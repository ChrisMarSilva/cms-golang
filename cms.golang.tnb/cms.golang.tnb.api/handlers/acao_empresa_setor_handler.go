package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type AcaoEmpresaSetorHandler struct {
	service service.AcaoEmpresaSetorService
}

func NewAcaoEmpresaSetorHandler(service service.AcaoEmpresaSetorService) *AcaoEmpresaSetorHandler {
	return &AcaoEmpresaSetorHandler{
		service: service,
	}
}

func (h *AcaoEmpresaSetorHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
