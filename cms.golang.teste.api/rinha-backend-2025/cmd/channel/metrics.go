package main

// import (
// 	"net/http"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// var (
// 	jobsIngress = prometheus.NewCounter(prometheus.CounterOpts{
// 		Name: "jobs_ingress_total",
// 		Help: "Total de jobs enfileirados",
// 	})
// 	jobsFallback = prometheus.NewCounter(prometheus.CounterOpts{
// 		Name: "jobs_fallback_total",
// 		Help: "Total de jobs que caíram no fallback",
// 	})
// 	queueLength = prometheus.NewGauge(prometheus.GaugeOpts{
// 		Name: "jobs_queue_length",
// 		Help: "Tamanho atual da fila de jobs",
// 	})
// )

func initMetrics() {
	// 	prometheus.MustRegister(jobsIngress, jobsFallback, queueLength)
}

func metricsServer() {
	// 	ticker := time.NewTicker(500 * time.Millisecond)
	// 	defer ticker.Stop()

	// 	http.Handle("/metrics", promhttp.Handler())
	// 	go func() { http.ListenAndServe(":2112", nil) }()

	// 	for range ticker.C {
	// 		// atualiza tamanho da fila
	// 		queueLength.Set(float64(capMetricsQueue()))
	// 	}
}
