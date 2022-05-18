package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type AcaoEmpresaSegmentoHandler struct {
	service service.AcaoEmpresaSegmentoService
}

func NewAcaoEmpresaSegmentoHandler(service service.AcaoEmpresaSegmentoService) *AcaoEmpresaSegmentoHandler {
	return &AcaoEmpresaSegmentoHandler{
		service: service,
	}
}

func (h *AcaoEmpresaSegmentoHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
