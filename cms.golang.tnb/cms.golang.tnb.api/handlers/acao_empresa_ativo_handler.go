package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type AcaoEmpresaAtivoHandler struct {
	service service.AcaoEmpresaAtivoService
}

func NewAcaoEmpresaAtivoHandler(service service.AcaoEmpresaAtivoService) *AcaoEmpresaAtivoHandler {
	return &AcaoEmpresaAtivoHandler{
		service: service,
	}
}

func (h *AcaoEmpresaAtivoHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AcaoEmpresaAtivoHandler) GetListaCodigoCompleto(c *gin.Context) {

	list, err := h.service.GetListaCodigoCompleto()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AcaoEmpresaAtivoHandler) GetListaCodigoCompletoAcao(c *gin.Context) {

	list, err := h.service.GetListaCodigoCompletoAcao()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AcaoEmpresaAtivoHandler) GetListaCodigoCompletoFii(c *gin.Context) {

	list, err := h.service.GetListaCodigoCompletoFii()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AcaoEmpresaAtivoHandler) GetListaCodigoCompletoEtf(c *gin.Context) {

	list, err := h.service.GetListaCodigoCompletoEtf()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AcaoEmpresaAtivoHandler) GetListaCodigoCompletoBrd(c *gin.Context) {

	list, err := h.service.GetListaCodigoCompletoBrd()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AcaoEmpresaAtivoHandler) GetListaCodigoCompletoCripto(c *gin.Context) {

	list, err := h.service.GetListaCodigoCompletoCripto()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
