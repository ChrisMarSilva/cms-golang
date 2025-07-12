package handlers

import (
	"net/http"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/internals/usecases"
	"github.com/chrismarsilva/rinha-backend-2025/internals/utils"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	//"github.com/prometheus/client_golang/prometheus"
)

type Handlers struct {
	useCases *usecases.UseCases
	config   *utils.Config
}

// var (
//     ReqCounter   = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "rinha", Name: "http_requests_total", Help: "Total HTTP requests"}, []string{"path", "method", "status"})
//     JobQueueLen  = prometheus.NewGauge(prometheus.GaugeOpts{Namespace: "rinha", Name: "job_queue_length", Help: "Length of payment queue"})
// )

// func init() {
//     prometheus.MustRegister(ReqCounter, JobQueueLen)
// }

func New(useCases *usecases.UseCases, config *utils.Config) *Handlers {
	return &Handlers{
		useCases: useCases,
		config:   config,
	}
}

func (h Handlers) Listen() error {
	// gin.SetMode(cfg.GinMode)
	// gin.SetMode(gin.ReleaseMode) // DebugMode // ReleaseMode
	//gin.DisableConsoleColor()
	// gin.ForceConsoleColor()

	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"::1", "127.0.0.1", "192.168.1.2", "10.0.0.0/8"})

	router.Use(gin.Recovery())
	router.Use(ErrorHandler())
	router.Use(gzip.Gzip(gzip.BestSpeed)) // DefaultCompression
	// router.Use(infra.PrometheusMiddleware())

	router.GET("/", GetDefault)
	//router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// router.POST("/payments", h.Pay);
	// router.GET("/summary", h.Summary)

	h.registerUserEndpoints(router)

	s := &http.Server{
		Addr:         h.config.UriPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	return s.ListenAndServe()
	//return router.Run(config.UriPort)
}

func GetDefault(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// func PrometheusMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Next()
//         ReqCounter.WithLabelValues(c.FullPath(), c.Request.Method, fmt.Sprint(c.Writer.Status())).Inc()
//     }
// }

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			//c.JSON(http.StatusInternalServerError, map[string]any{"success": false, "message": err.Error()})
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		}
	}
}
