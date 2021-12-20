package main

import (
	"github.com/TaranovDmitry/Microservices"
	"github.com/TaranovDmitry/Microservices/pkg/handler"
	"github.com/TaranovDmitry/Microservices/repository"
	"github.com/TaranovDmitry/Microservices/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialzing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: "localhost",
		Port: "5432",
		Username: "postgres",
		Password: "12345",
		DBName: "todo-db",
		SSLMode: "disable",
	})
	if err != nil{
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Microservices.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouts()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
