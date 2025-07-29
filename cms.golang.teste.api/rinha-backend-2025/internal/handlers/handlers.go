package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/internal/adapters"
	"github.com/chrismarsilva/rinha-backend-2025/internal/dtos"
	"github.com/chrismarsilva/rinha-backend-2025/internal/services"
	"github.com/chrismarsilva/rinha-backend-2025/internal/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define metrics
var (
	HttpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_total",
		Help: "Total number of requests processed by the API",
	}, []string{"path", "status"})

	HttpRequestErrorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_error_total",
		Help: "Total number of errors returned by the API",
	}, []string{"path", "status"})

	HttpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

//var customRegistry = prometheus.NewRegistry()

func init() {
	//customRegistry.MustRegister(HttpRequestTotal, HttpRequestErrorTotal, HttpRequestDuration)
	prometheus.Register(HttpRequestTotal)
	prometheus.Register(HttpRequestErrorTotal)
	prometheus.Register(HttpRequestDuration)
}

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
	router.Use(h.ErrorsMiddleware())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(h.RequestMetricsMiddleware())

	router.POST("/payments", h.PaymentHandler)
	router.GET("/payments-summary", h.SummaryHandler)
	router.GET("/health", h.HealthHandler)
	//router.GET("/metrics", h.PrometheusHandler())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", h.config.UriPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s.ListenAndServe()
}

func (h Handlers) ErrorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		}
	}
}

func (h Handlers) RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Writer.Status()
		path := c.Request.URL.Path

		timer := prometheus.NewTimer(HttpRequestDuration.WithLabelValues(path))
		c.Next()
		timer.ObserveDuration()

		HttpRequestTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		if status >= 400 {
			HttpRequestErrorTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		}
	}
}

func (h Handlers) PrometheusHandler() gin.HandlerFunc {
	//handler := promhttp.HandlerFor(customRegistry, promhttp.HandlerOpts{})
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
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
