package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ConfigLoader struct {
	v          *viper.Viper
	configName string
	configType string
	configPath []string
	env        string
}

type IDBConfiguration interface {
	ConnectDatabase() (*gorm.DB, error)
}
