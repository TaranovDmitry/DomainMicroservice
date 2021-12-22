package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/TaranovDmitry/Microservices/entity"
)

type PortService interface {
	AllPorts() (entity.Ports, error)
	Update(ports entity.Ports) error
}

type Handler struct {
	service PortService
}

func NewHandler(services PortService) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	api.GET("/ports", h.allPorts)
	api.POST("/ports", h.updateList)

	return router
}
