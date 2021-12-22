package main

import (
	"context"
	"github.com/TaranovDmitry/Microservices/config"
	"github.com/TaranovDmitry/Microservices/handlers"
	"github.com/TaranovDmitry/Microservices/repository"
	"github.com/TaranovDmitry/Microservices/services"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialzing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg := config.New()

	db, err := repository.NewPostgresDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	portsRepository := repository.NewPortsRepository(db)
	service := services.NewService(portsRepository)
	handler := handlers.NewHandler(service)

	var srv Server
	if err := srv.Run(cfg.Port, handler.InitRouts()); err != nil {
		logrus.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
