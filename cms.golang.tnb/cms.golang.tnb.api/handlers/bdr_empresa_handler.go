package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type BdrEmpresaHandler struct {
	service service.BdrEmpresaService
}

func NewBdrEmpresaHandler(service service.BdrEmpresaService) *BdrEmpresaHandler {
	return &BdrEmpresaHandler{
		service: service,
	}
}

func (h *BdrEmpresaHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
