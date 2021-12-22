package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
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

func New() (Config, error) {
	if err := initConfig(); err != nil {
		return Config{}, fmt.Errorf("failed to init config: %w", err)
	}

	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to get env var: %w", err)
	}

	return Config{
		Port: viper.GetString("port"),
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		}}, nil
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
