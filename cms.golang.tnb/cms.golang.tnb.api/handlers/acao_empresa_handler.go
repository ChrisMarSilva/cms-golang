package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type AcaoEmpresaHandler struct {
	service service.AcaoEmpresaService
}

func NewAcaoEmpresaHandler(service service.AcaoEmpresaService) *AcaoEmpresaHandler {
	return &AcaoEmpresaHandler{
		service: service,
	}
}

func (h *AcaoEmpresaHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
