package handlers

import (
	"net/http"

	"github.com/chrismarsilva/rinha-backend-2025/internals/dtos"
	"github.com/gin-gonic/gin"
)

func (h Handlers) registerUserEndpoints(router *gin.Engine) {
	routerUsers := router.Group("/v1/users")
	routerUsers.GET("/", h.getAllUsers)
	routerUsers.POST("/", h.createUser)
}

func (h Handlers) getAllUsers(c *gin.Context) {
	users, err := h.useCases.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := gin.H{"message": "ok", "qtd": len(users), "users": users}
	c.JSON(http.StatusOK, data)
}

func (h Handlers) createUser(c *gin.Context) {
	var request dtos.UserRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.useCases.Create(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//c.Error(err)
		return
	}

	data := gin.H{"message": "ok", "user": gin.H{"id": user.ID, "name": user.Name, "email": user.Email}}
	c.JSON(http.StatusCreated, data)
}
