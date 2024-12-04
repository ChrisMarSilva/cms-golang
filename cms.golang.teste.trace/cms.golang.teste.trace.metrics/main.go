package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.trace.metrics
// go get github.com/prometheus/client_golang/prometheus
// go get github.com/prometheus/client_golang/prometheus/promauto
// go get github.com/prometheus/client_golang/prometheus/promhttp
// go mod tidy

// go run main.go

// docker-compose down
// docker-compose up -d --build

// ab -n 100000 -c 10 http://localhost:8080/
// k6 run ./k6/index.js
// k6 run --vus 1000 --duration 1m ./k6/index.js

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cms_http_request_total", 
			Help: "Number of requests",
		}, []string{"path"}, // []string{"user", "status"}, // labels
	)
	
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "cms_http_request_duration_seconds", 
			Help: "Duration of HTTP requests",
		}, []string{"path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(httpDuration)
}


func handle(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(r.URL.Path))
	defer timer.ObserveDuration()

	defer func() {
		httpRequests.WithLabelValues(r.URL.Path).Inc()
		// httpRequests.WithLabelValues(user, status).Inc()
	}()

	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", handle) //http.Handle("/", http.HandlerFunc(handle))
	http.Handle("/metrics", promhttp.Handler())
	
	log.Println("Application up and running")
	http.ListenAndServe(":8080", nil)
}
