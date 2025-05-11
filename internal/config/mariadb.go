package config

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type newDBConfiguration struct {
	v *viper.Viper
}

func NewMariaDBConfiguration(v *viper.Viper) IDBConfiguration {
	return &newDBConfiguration{
		v: v,
	}
}

func (n *newDBConfiguration) ConnectDatabase() (*gorm.DB, error) {
	dbViperSection := n.v.Sub("database")
	if dbViperSection == nil {
		log.Fatalf("Bagian konfigurasi 'database' tidak ditemukan dalam file konfigurasi.")
	}

	dsn := dbViperSection.GetString("username") + ":" + dbViperSection.GetString("password") + "@tcp(" + dbViperSection.GetString("host") + ":" + dbViperSection.GetString("port") + ")/" + dbViperSection.GetString("database_name") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbConn, err := db.DB()
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxIdleConns(dbViperSection.GetInt("db_max_idle_conns"))
	dbConn.SetMaxOpenConns(n.v.GetInt("db_max_open_conns"))
	dbConn.SetConnMaxLifetime(n.v.GetDuration("db_conn_max_lifetime"))
	dbConn.SetConnMaxIdleTime(n.v.GetDuration("db_conn_max_idle_time"))

	return db, nil
}
