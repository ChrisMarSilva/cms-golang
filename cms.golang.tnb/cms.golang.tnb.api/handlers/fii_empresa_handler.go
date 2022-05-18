package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type FiiEmpresaHandler struct {
	service service.FiiEmpresaService
}

func NewFiiEmpresaHandler(service service.FiiEmpresaService) *FiiEmpresaHandler {
	return &FiiEmpresaHandler{
		service: service,
	}
}

func (h *FiiEmpresaHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
