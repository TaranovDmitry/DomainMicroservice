package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	Username string
	SSLMode  string
	Password string
}

func New() Config {
	return Config{
		Port: viper.GetString("port"),
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		}}
}
