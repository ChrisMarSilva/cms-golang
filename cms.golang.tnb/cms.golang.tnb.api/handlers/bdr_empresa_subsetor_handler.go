package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type BdrEmpresaSubSetorHandler struct {
	service service.BdrEmpresaSubSetorService
}

func NewBdrEmpresaSubSetorHandler(service service.BdrEmpresaSubSetorService) *BdrEmpresaSubSetorHandler {
	return &BdrEmpresaSubSetorHandler{
		service: service,
	}
}

func (h *BdrEmpresaSubSetorHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
