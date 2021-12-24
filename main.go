package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/TaranovDmitry/DomainMicroservice/config"
	"github.com/TaranovDmitry/DomainMicroservice/handlers"
	"github.com/TaranovDmitry/DomainMicroservice/repository"
	"github.com/TaranovDmitry/DomainMicroservice/services"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetOutput(os.Stdout)

	cfg, err := config.New()
	if err != nil {
		logger.Fatalf("failed to initialize config %v", err)
	}

	db, err := repository.NewPostgresDB(cfg.DB)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	portsRepository := repository.NewPortsRepository(db)
	service := services.NewService(portsRepository)
	handler := handlers.NewHandler(service)

	var srv Server
	if err := srv.Run(cfg.Port, handler.InitRouts()); err != nil {
		_ = srv.Shutdown(context.TODO())
		logger.Fatalf("error occured while running http server %s", err.Error())
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
