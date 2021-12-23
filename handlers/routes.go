package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/TaranovDmitry/DomainMicroservice/entity"
)

type PortService interface {
	AllPorts() (entity.Ports, error)
	Upsert(ports entity.Ports) error
}

type Handler struct {
	service PortService
}

func NewHandler(services PortService) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	domain := router.Group("/domain")
	domainV1 := domain.Group("/v1")
	domainV1.GET("/ports", h.allPorts)
	domainV1.POST("/ports", h.updateList)

	return router
}
