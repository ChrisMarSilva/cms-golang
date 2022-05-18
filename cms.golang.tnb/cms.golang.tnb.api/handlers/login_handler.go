package handler

import (
	"net/http"

	service "github.com/ChrisMarSilva/cms.golang.tnb.api/services"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service service.LoginService
}

func NewLoginHandler(service service.LoginService) *LoginHandler {
	return &LoginHandler{
		service: service,
	}
}

func (h *LoginHandler) Entrar(c *gin.Context) {

	txtEmail := c.Query("txtEmail")
	txtSenha := c.Query("txtSenha")

	if txtEmail == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "E-mail não informado!"})
		return
	}

	if txtSenha == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Senha não informada!"})
		return
	}

	row, err := h.service.Entrar(txtEmail, txtSenha)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": gin.H{
				"Resultado": "ok",
				"Mensagem":  "",
				"Dados": gin.H{
					"Id":    row.ID,
					"Tipo":  row.Tipo,
					"Nome":  row.Nome,
					"Foto":  row.Foto,
					"Email": row.Email,
				},
			},
		},
	)
}
