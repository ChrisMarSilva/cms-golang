package handlers

import (
	"net/http"

	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/services"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PersonHandler struct {
	Service *services.PersonService
}

func NewPersonHandler(service *services.PersonService) *PersonHandler {
	return &PersonHandler{
		Service: service,
	}
}

func (h PersonHandler) Add(c *gin.Context) {
	ctx := c.Request.Context()

	ctx, span := utils.Tracer.Start(ctx, "PersonHandler.Add")
	defer span.End()

	var request dtos.PersonRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error-invalid JSON body": err.Error()})
		return
	}

	err := h.Service.Add(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add person: " + err.Error()})
		return
	}

	//c.JSON(http.StatusCreated, response)
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (h PersonHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()

	ctx, span := utils.Tracer.Start(ctx, "PersonHandler.GetAll")
	defer span.End()

	persons, err := h.Service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get all persons: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (h PersonHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	ctx, span := utils.Tracer.Start(ctx, "PersonHandler.GetByID")
	defer span.End()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	person, err := h.Service.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found person by ID: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h PersonHandler) GetCount(c *gin.Context) {
	ctx := c.Request.Context()

	ctx, span := utils.Tracer.Start(ctx, "PersonHandler.GetCount")
	defer span.End()

	response, err := h.Service.GetCount(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found person by ID: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
