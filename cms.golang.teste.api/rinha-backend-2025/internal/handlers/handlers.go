package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/internal/adapters"
	"github.com/chrismarsilva/rinha-backend-2025/internal/dtos"
	"github.com/chrismarsilva/rinha-backend-2025/internal/services"
	"github.com/chrismarsilva/rinha-backend-2025/internal/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	config     *utils.Config
	PaymentSvc *services.PaymentService
	SummarySvc *services.SummaryService
}

func New(config *utils.Config, paymentSvc *services.PaymentService, summarySvc *services.SummaryService) *Handlers {
	return &Handlers{
		config:     config,
		PaymentSvc: paymentSvc,
		SummarySvc: summarySvc,
	}
}

func (h Handlers) Listen() error {
	gin.SetMode(h.config.GinMode)

	router := gin.Default()
	router.Use(h.ErrorsHandler())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.BestSpeed))

	router.POST("/payments", h.PaymentHandler)
	router.GET("/payments-summary", h.SummaryHandler)
	router.GET("/health", h.HealthHandler)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", h.config.UriPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s.ListenAndServe()
}

func (h Handlers) ErrorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		}
	}
}

func (h Handlers) PaymentHandler(c *gin.Context) {
	var request dtos.PaymentRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error-invalid JSON body": err.Error()})
		return
	}

	if request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error-invalid amount": "amount must be greater than 0"})
		return
	}

	adapters.Jobs <- request

	response := gin.H{"status": "success", "message": "Payment request accepted"} // , "correlationId": request.CorrelationId
	c.JSON(http.StatusAccepted, response)
}

func (h Handlers) SummaryHandler(c *gin.Context) {
	fromStr := c.Query("from")
	toStr := c.Query("to")
	var from, to *time.Time

	if fromStr != "" {
		parsedFrom, err := time.Parse(time.RFC3339, fromStr)
		if err == nil {
			from = &parsedFrom
		}
	}

	if toStr != "" {
		parsedTo, err := time.Parse(time.RFC3339, toStr)
		if err == nil {
			to = &parsedTo
		}
	}

	summary, err := h.SummarySvc.GetSummary(c.Request.Context(), from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get summary: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}

func (h Handlers) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
