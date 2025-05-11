package inbound

import (
	"context"

	"github.com/fanchann/grpc-crud/internal/core/domain/dto"
)

type IBookServicePorts interface {
	CreateNewBook(ctx context.Context, req *dto.CreateBookRequest) error
	GetAllBooks(ctx context.Context) []dto.BookResponse
	DeleteBook(ctx context.Context, id string) error
	UpdateBook(ctx context.Context, id string, req *dto.UpdateBookRequest) error
	GetBookByID(ctx context.Context, id string) (*dto.BookResponse, error)
}
