package handlers

import (
	"github.com/TaranovDmitry/Microservices/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	api.GET("/ports", h.allPorts)
	api.POST("/ports", h.updateList)

	return router
}
