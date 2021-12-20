package handler

import (
	"github.com/TaranovDmitry/Microservices/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	lists := router.Group("/api")
	{
		lists.GET("/", h.getAllLists)
		lists.POST("/", h.updateList)
	}
	return router
}