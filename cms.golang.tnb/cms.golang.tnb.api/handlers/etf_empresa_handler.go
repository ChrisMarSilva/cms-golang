package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type EtfEmpresaHandler struct {
	service service.EtfEmpresaService
}

func NewEtfEmpresaHandler(service service.EtfEmpresaService) *EtfEmpresaHandler {
	return &EtfEmpresaHandler{
		service: service,
	}
}

func (h *EtfEmpresaHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
