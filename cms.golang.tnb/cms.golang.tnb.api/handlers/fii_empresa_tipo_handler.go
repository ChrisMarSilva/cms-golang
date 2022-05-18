package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type FiiEmpresaTipoHandler struct {
	service service.FiiEmpresaTipoService
}

func NewFiiEmpresaTipoHandler(service service.FiiEmpresaTipoService) *FiiEmpresaTipoHandler {
	return &FiiEmpresaTipoHandler{
		service: service,
	}
}

func (h *FiiEmpresaTipoHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
