package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type BdrEmpresaSegmentoHandler struct {
	service service.BdrEmpresaSegmentoService
}

func NewBdrEmpresaSegmentoHandler(service service.BdrEmpresaSegmentoService) *BdrEmpresaSegmentoHandler {
	return &BdrEmpresaSegmentoHandler{
		service: service,
	}
}

func (h *BdrEmpresaSegmentoHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
