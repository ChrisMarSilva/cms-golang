package handlers

import (
	"log/slog"
	"net/http"

	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/services"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PersonHandler struct {
	logger  *slog.Logger
	service *services.PersonService
}

func NewPersonHandler(logger *slog.Logger, service *services.PersonService) *PersonHandler {
	return &PersonHandler{
		logger:  logger,
		service: service,
	}
}

func (h PersonHandler) Add(c *gin.Context) {
	ctx, span := utils.Tracer.Start(c.Request.Context(), "PersonHandler.Add")
	defer span.End()

	var request dtos.PersonRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Add(ctx, request)
	if err != nil {
		h.logger.Error("Failed to add person", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h PersonHandler) GetAll(c *gin.Context) {
	ctx, span := utils.Tracer.Start(c.Request.Context(), "PersonHandler.GetAll")
	defer span.End()

	persons, err := h.service.GetAll(ctx)
	if err != nil {
		h.logger.Error("Failed to get all persons", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}

func (h PersonHandler) GetByID(c *gin.Context) {
	ctx, span := utils.Tracer.Start(c.Request.Context(), "PersonHandler.GetByID")
	defer span.End()

	idStr := c.Param("id")
	h.logger.Info("Getting person by ID", slog.Any("id", idStr))

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("Failed to parse ID", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := h.service.GetByID(ctx, id)
	if err != nil {
		h.logger.Error("Failed to get person by ID", slog.Any("error", err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h PersonHandler) GetCount(c *gin.Context) {
	ctx, span := utils.Tracer.Start(c.Request.Context(), "PersonHandler.GetCount")
	defer span.End()

	response, err := h.service.GetCount(ctx)
	if err != nil {
		h.logger.Error("Failed to get person count", slog.Any("error", err))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h PersonHandler) Update(c *gin.Context) {
	ctx, span := utils.Tracer.Start(c.Request.Context(), "PersonHandler.Update")
	defer span.End()

	idStr := c.Param("id")
	h.logger.Info("Getting person by ID", slog.Any("id", idStr))

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("Failed to parse ID", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request dtos.PersonRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// exists, err := h.service.ExistByID(ctx, id)
	// if err != nil {
	// 	h.logger.Error("Failed to get person by ID", slog.Any("error", err))
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }

	// if !exists {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	// 	return
	// }

	err = h.service.Update(ctx, id, request)
	if err != nil {
		h.logger.Error("Failed to update person", slog.Any("error", err))
		if err.Error() == "Not found." {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (h PersonHandler) Delete(c *gin.Context) {
	ctx, span := utils.Tracer.Start(c.Request.Context(), "PersonHandler.Delete")
	defer span.End()

	idStr := c.Param("id")
	h.logger.Info("Getting person by ID", slog.Any("id", idStr))

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.logger.Error("Failed to parse ID", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// exists, err := h.service.ExistByID(ctx, id)
	// if err != nil {
	// 	h.logger.Error("Failed to get person by ID", slog.Any("error", err))
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }

	// if !exists {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	// 	return
	// }

	err = h.service.DeleteByID(ctx, id)
	if err != nil {
		h.logger.Error("Failed to delete person", slog.Any("error", err))
		if err.Error() == "Not found." {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
