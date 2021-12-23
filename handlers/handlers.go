package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

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

type err struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, err{message})
}

func (h *Handler) allPorts(c *gin.Context) {
	ports, err := h.service.AllPorts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ports)
}

func (h *Handler) updateList(c *gin.Context) {
	var ports entity.Ports

	err := c.BindJSON(&ports)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Upsert(ports)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
