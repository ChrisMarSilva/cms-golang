package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type BdrEmpresaSetorHandler struct {
	service service.BdrEmpresaSetorService
}

func NewBdrEmpresaSetorHandler(service service.BdrEmpresaSetorService) *BdrEmpresaSetorHandler {
	return &BdrEmpresaSetorHandler{
		service: service,
	}
}

func (h *BdrEmpresaSetorHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
