package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type FiiEmpresaAdminHandler struct {
	service service.FiiEmpresaAdminService
}

func NewFiiEmpresaAdminHandler(service service.FiiEmpresaAdminService) *FiiEmpresaAdminHandler {
	return &FiiEmpresaAdminHandler{
		service: service,
	}
}

func (h *FiiEmpresaAdminHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
