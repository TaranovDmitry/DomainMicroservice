package main

import (
	"context"
	"net/http"
	"time"

	"github.com/TaranovDmitry/Microservices/config"
	"github.com/TaranovDmitry/Microservices/handlers"
	"github.com/TaranovDmitry/Microservices/repository"
	"github.com/TaranovDmitry/Microservices/services"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := config.New()
	if err != nil {
		logrus.Fatalf("failed to initialize config %v", err)
	}

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

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
