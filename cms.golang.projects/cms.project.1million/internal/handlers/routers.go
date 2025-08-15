package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
)

type Router struct {
	Config        *utils.Config
	PersonHandler *PersonHandler
}

func NewRouter(config *utils.Config, personHandler *PersonHandler) *Router {
	return &Router{
		Config:        config,
		PersonHandler: personHandler,
	}
}

// func TrackMetrics() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		path := c.Request.URL.Path
// 		c.Next()

// 		status := c.Writer.Status()
// 		utils.MetricHttpRequestCount.WithLabelValues(path, http.StatusText(status)).Inc()

// 		if status >= http.StatusBadRequest {
// 			utils.MetricHttpErrorCount.WithLabelValues(path, http.StatusText(status)).Inc()
// 		}
// 	}
// }

func prometheusMiddleware(c *gin.Context) {
	path := c.Request.URL.Path
	timer := prometheus.NewTimer(utils.MetricHttpRequestDuration.WithLabelValues(path)) // begin timer to measure the requests duration

	utils.MetricHttpRequestsTotal.WithLabelValues(path).Inc() // increment total request counter
	utils.MetricHttpActiveConnections.Inc()                   // increment number of active connections

	c.Next() // complete processing request

	timer.ObserveDuration()                 // record request duration (post processing)
	utils.MetricHttpActiveConnections.Dec() // decrement total number of active connections (post processing)
}

func (r Router) Listen() error {
	gin.SetMode(r.Config.GinMode)
	if r.Config.GinMode == gin.ReleaseMode || r.Config.GinMode == gin.TestMode {
		gin.DefaultWriter = io.Discard // Disable default logger in release mode
	}

	prometheus.MustRegister(utils.MetricHttpRequestsTotal)
	prometheus.MustRegister(utils.MetricHttpRequestDuration)
	prometheus.MustRegister(utils.MetricHttpActiveConnections)
	// prometheus.MustRegister(utils.MetricHttpCounter)
	// prometheus.MustRegister(utils.MetricHttpRequestCount)
	// prometheus.MustRegister(utils.MetricHttpErrorCount)

	router := gin.Default()
	router.Use(r.ErrorsMiddleware())
	router.Use(gin.Recovery())
	//router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(otelgin.Middleware("cms.api.1million"))

	//router.Use(TrackMetrics())
	router.Use(prometheusMiddleware)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	//router.GET("/metrics", gin.WrapH(utils.PromHandler))
	// p := ginprometheus.NewPrometheus("")
	// p.Use(router)

	router.GET("/health", r.HealthHandler)

	router.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		//start := time.Now()

		ctx, span := utils.Tracer.Start(ctx, "INI")
		span.SetAttributes(attribute.Key("test-set-1").String("value-1"))
		span.AddEvent("Health check event-1")
		defer span.End()

		time.Sleep(500 * time.Millisecond)
		span.SetAttributes(attribute.Key("test-set-2").String("value-2"))
		span.AddEvent("Health check event-2")

		time.Sleep(500 * time.Millisecond)
		span.SetAttributes(attribute.Key("test-set-3").String("value-3"))
		span.AddEvent("Health check event-3")

		// // Métricas
		// value := rand.Float64() * 100 // simula valor da venda
		// utils.MetricHttpRequestCount.Add(c, 1)
		// utils.MetricHttpLatency.Record(c, time.Since(start).Seconds())
		// utils.MetricTotalSalesValue.Add(c, value)
		// utils.MetricPingRequestCount.Add(c, 1)
		//counter, _ := utils.Meter.Float64Counter("foo", api.WithDescription("a simple counter"))
		//counter.Add(ctx, 5)
		// utils.MetricHttpCounter.WithLabelValues("AAA").Inc()
		// utils.MetricHttpCounter.WithLabelValues("BBB").Inc()
		// utils.MetricHttpCounter.WithLabelValues("BBB").Inc()

		c.JSON(http.StatusOK, "ok")
	})

	v1Group := router.Group("/v1")
	{
		personGroup := v1Group.Group("/person")
		{
			personGroup.POST("", r.PersonHandler.Add)
			personGroup.GET("/", r.PersonHandler.GetAll)
			personGroup.GET("/:id", r.PersonHandler.GetByID)
			personGroup.GET("/count", r.PersonHandler.GetCount)
		}
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", r.Config.UriPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s.ListenAndServe()
}

func (r Router) ErrorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		}
	}
}

func (r Router) HealthHandler(c *gin.Context) {
	//c.JSON(http.StatusOK, r.Config.NameApi)
	c.JSON(http.StatusOK, "ok")
}
