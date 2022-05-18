package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type AlertaAssinaturaHandler struct {
	service service.AlertaAssinaturaService
}

func NewAlertaAssinaturaHandler(service service.AlertaAssinaturaService) *AlertaAssinaturaHandler {
	return &AlertaAssinaturaHandler{
		service: service,
	}
}

func (h *AlertaAssinaturaHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
