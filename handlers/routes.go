package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	domain := router.Group("/domain")
	domainV1 := domain.Group("/v1")
	domainV1.GET("/ports", h.allPorts)
	domainV1.POST("/ports", h.updateList)

	return router
}
