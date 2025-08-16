package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	//sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func init() {
	// prometheus.MustRegister(utils.MetricHttpRequestsTotal)
	// prometheus.MustRegister(utils.MetricHttpRequestDuration)
	// prometheus.MustRegister(utils.MetricHttpActiveConnections)
}

type Router struct {
	logger  *slog.Logger
	config  *utils.Config
	handler *PersonHandler
	rdb     *stores.RedisCache
}

func NewRouter(logger *slog.Logger, config *utils.Config, handler *PersonHandler, rdb *stores.RedisCache) *Router {
	return &Router{
		logger:  logger,
		config:  config,
		handler: handler,
		rdb:     rdb,
	}
}

func (r Router) Listen() error {
	gin.SetMode(r.config.GinMode)
	if r.config.GinMode == gin.ReleaseMode || r.config.GinMode == gin.TestMode {
		gin.DefaultWriter = io.Discard
	} else {
		gin.ForceConsoleColor()
	}

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.BestSpeed, gzip.WithExcludedPaths([]string{"/metrics"})))
	router.Use(gin.Recovery())
	router.Use(r.ErrorsMiddleware())
	router.Use(r.IdempotencyMiddleware())
	// router.Use(r.PrometheusMiddleware())
	router.Use(otelgin.Middleware("cms.api.1million"))
	//router.Use(sloggin.New(r.logger))

	router.GET("/", r.HomeHandler)
	router.GET("/ping", r.PingHandler)
	router.GET("/health", r.HealthHandler)
	//router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1Group := router.Group("/v1")
	{
		personGroup := v1Group.Group("/person")
		{
			personGroup.POST("", r.handler.Add)
			personGroup.GET("", r.handler.GetAll)
			personGroup.GET("/:id", r.handler.GetByID)
			personGroup.PUT("/:id", r.handler.Update)
			personGroup.DELETE("/:id", r.handler.Delete)
			personGroup.GET("/count", r.handler.GetCount)
		}
	}

	r.logger.Info("API rodando na porta", slog.String("port", r.config.UriPort))

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", r.config.UriPort),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return s.ListenAndServe()
}

// func (r Router) PrometheusMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		path := c.Request.URL.Path
// 		timer := prometheus.NewTimer(utils.MetricHttpRequestDuration.WithLabelValues(path)) // begin timer to measure the requests duration

// 		utils.MetricHttpRequestsTotal.WithLabelValues(path).Inc() // increment total request counter
// 		utils.MetricHttpActiveConnections.Inc()                   // increment number of active connections

// 		// 		status := c.Writer.Status()
// 		// 		utils.MetricHttpRequestCount.WithLabelValues(path, http.StatusText(status)).Inc()

// 		// 		if status >= http.StatusBadRequest {
// 		// 			utils.MetricHttpErrorCount.WithLabelValues(path, http.StatusText(status)).Inc()
// 		// 		}

// 		c.Next() // complete processing request

// 		timer.ObserveDuration()                 // record request duration (post processing)
// 		utils.MetricHttpActiveConnections.Dec() // decrement total number of active connections (post processing)
// 	}
// }

func (r Router) ErrorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// r.logger.Info("Request processed successfully", slog.Int("errors", len(c.Errors)))
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			r.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	r.status = r.ResponseWriter.Status()
	return r.ResponseWriter.Write(b)
}

type ResponseCache struct {
	Status int             `json:"status"`
	Header http.Header     `json:"header"`
	Data   json.RawMessage `json:"data"`
}

func NewResponseCache(status int, header http.Header, data []byte) ResponseCache {
	return ResponseCache{
		Status: status,
		Header: header,
		Data:   data,
	}
}

func (r Router) IdempotencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			c.Next()
			return
		}

		key := c.GetHeader("X-Idempotency-Key")
		if key == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "X-Idempotency-Key header is required"})
			c.Abort()
			return
		}

		payloadGet, err := r.rdb.HGet(c.Request.Context(), "idempotency", key)
		if err == nil {
			var dataGet ResponseCache
			err := sonic.Unmarshal([]byte(payloadGet), &dataGet)
			if err != nil {
				r.logger.Error("Failed to unmarshal cache data", slog.Any("error", err))
				return
			}

			c.Header("X-Idempotency-Key", key)
			c.Header("X-Result", "Redis")

			for k, vals := range dataGet.Header {
				for _, v := range vals {
					if k == "X-Result" {
						c.Writer.Header().Set(k, "Redis")
					} else {
						c.Writer.Header().Add(k, v)
					}
				}
			}

			c.JSON(dataGet.Status, dataGet.Data)
			c.Abort()
			return
		}

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		c.Header("X-Idempotency-Key", key)
		c.Header("X-Result", "Normal")
		c.Next()

		dataSet := NewResponseCache(c.Writer.Status(), c.Writer.Header(), w.body.Bytes())

		payloadSet, err := sonic.Marshal(dataSet)
		if err != nil {
			r.logger.Error("Failed to marshal response cache", slog.Any("error", err))
			return
		}

		err = r.rdb.HSet(c.Request.Context(), "idempotency", key, payloadSet)
		if err != nil {
			r.logger.Error("Failed to set idempotency key", slog.Any("error", err))
			return
		}
	}
}

func (r Router) HomeHandler(c *gin.Context) {
	r.logger.Info("Processing Home Page request")

	ctx, span1 := utils.Tracer.Start(c.Request.Context(), "span1")
	defer span1.End()

	time.Sleep(500 * time.Millisecond)

	ctx, span2 := utils.Tracer.Start(ctx, "span2")
	defer span2.End()

	time.Sleep(500 * time.Millisecond)

	ctx, span3 := utils.Tracer.Start(ctx, "span3")
	defer span3.End()

	time.Sleep(1000 * time.Millisecond)

	r.logger.Info("1-Processing example use case")
	utils.Logger.InfoContext(ctx, "2-Processing example use case")

	// // Métricas
	// value := rand.Float64() * 100 // simula valor da venda
	// utils.MetricHttpRequestCount.Add(c, 1)
	// utils.MetricHttpLatency.Record(c, time.Since(start).Seconds())
	// utils.MetricTotalSalesValue.Add(c, value)
	// utils.MetricPingRequestCount.Add(c, 1)
	// counter, _ := utils.Meter.Float64Counter("foo", api.WithDescription("a simple counter"))
	// counter.Add(ctx, 5)
	// utils.MetricHttpCounter.WithLabelValues("AAA").Inc()
	// utils.MetricHttpCounter.WithLabelValues("BBB").Inc()
	// utils.MetricHttpCounter.WithLabelValues("BBB").Inc()

	c.JSON(http.StatusOK, gin.H{"msg": "welcome"})
}

func (r Router) PingHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"msg": "pong"})
}

func (r Router) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
