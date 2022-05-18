package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type CorretoraListaHandler struct {
	service service.CorretoraListaService
}

func NewCorretoraListaHandler(service service.CorretoraListaService) *CorretoraListaHandler {
	return &CorretoraListaHandler{
		service: service,
	}
}

func (h *CorretoraListaHandler) GetLista(c *gin.Context) {

	list, err := h.service.GetLista()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
