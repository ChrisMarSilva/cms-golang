package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type CriptoEmpresaSituacaoHandler struct {
	service service.CriptoEmpresaSituacaoService
}

func NewCriptoEmpresaSituacaoHandler(service service.CriptoEmpresaSituacaoService) *CriptoEmpresaSituacaoHandler {
	return &CriptoEmpresaSituacaoHandler{
		service: service,
	}
}

func (h *CriptoEmpresaSituacaoHandler) GetSituacoes(c *gin.Context) {

	list, err := h.service.GetSituacoes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *CriptoEmpresaSituacaoHandler) GetSituacao(c *gin.Context) {

	codigo, _ := c.Params.Get("codigo")

	row, err := h.service.GetSituacao(codigo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": row})
}
