package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/go-playground/validator"
	"google.golang.org/grpc"

	grpcHandler "github.com/fanchann/grpc-crud/internal/adapter/grpc"
	"github.com/fanchann/grpc-crud/internal/adapter/storage"
	"github.com/fanchann/grpc-crud/internal/config"
	"github.com/fanchann/grpc-crud/internal/core/domain/entity"
	"github.com/fanchann/grpc-crud/internal/core/service"
	BooksProto "github.com/fanchann/grpc-crud/proto/protogen"
)

var configName *string

func init() {
	configName = flag.String("c", "config", "nama file konfigurasi")
	flag.Parse()

	if *configName == "" {
		log.Fatal("nama file konfigurasi tidak boleh kosong")
	}
}

func main() {
	cfgLoader := config.NewConfigLoader(*configName)

	fullConfig, err := cfgLoader.LoadConfig()
	if err != nil {
		log.Fatalf("gagal memuat konfigurasi: %v", err)
	}

	dbGormConfig := config.NewMariaDBConfiguration(fullConfig)

	mariadb, err := dbGormConfig.ConnectDatabase()
	if err != nil {
		log.Fatalf("gagal terhubung ke database: %v", err)
	}

	if err := mariadb.AutoMigrate(&entity.Books{}); err != nil {
		log.Fatalf("gagal melakukan AutoMigrate database: %v", err)
	}
	log.Println("automigrate selesai.")

	bookRepo := storage.NewBookRepoImpl(mariadb)
	validate := validator.New()
	bookService := service.NewBookService(bookRepo, validate)
	bookGrpcHandler := grpcHandler.NewBookGrpcImpl(bookService)

	serverPort := fullConfig.GetString("server.port")
	if serverPort == "" {
		log.Fatalf("port server (server.port) tidak dikonfigurasi.")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Fatalf("gagal melakukan listen: %v", err)
	}

	s := grpc.NewServer()
	BooksProto.RegisterBookServiceServer(s, bookGrpcHandler)

	log.Printf("server gRPC berjalan di %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("gagal menjalankan server gRPC: %v", err)
	}
}
