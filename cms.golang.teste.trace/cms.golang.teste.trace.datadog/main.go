package main

import (
	"log"
	"net/http"

	_ "github.com/pdrum/swagger-automation/docs"
	swaggerFiles "github.com/swaggo/files"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.trace.datadog
// go get gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer
// go get gopkg.in/DataDog/dd-trace-go.v1/profiler
// go get gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux
// go mod tidy

// swag init -g main.go
// swag init -g cmd/<package-name>/main.go
// swagger generate spec -o ./docs/swagger.json

// go run main.go

func main() {

	tracer.Start(tracer.WithServiceName("cms.golang.teste.trace.datadog"))
	defer tracer.Stop()

	http.HandleFunc("/", handler)
	http.HandleFunc("/home", handler2)
	http.HandleFunc("/index", handler3)
	http.HandleFunc("/swagger/*any", swaggerFiles.Handler)

	log.Fatal(http.ListenAndServe(":3001", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	span := tracer.StartSpan("web.request", tracer.ResourceName("/posts"))
	defer span.Finish()

	if span, ok := tracer.SpanFromContext(r.Context()); ok {
		span.SetTag("http.url", r.URL.Path)
	}

	span.SetTag("http.url", r.URL.Path)
	span.SetTag("<TAG_KEY>", "<TAG_VALUE>")

	w.Write([]byte("ok 1"))
}

func handler2(w http.ResponseWriter, r *http.Request) {

	if span, ok := tracer.SpanFromContext(r.Context()); ok {
		span.SetTag("http.url", r.URL.Path)
	}

	w.Write([]byte("ok 2"))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	span, ctx := tracer.StartSpanFromContext(r.Context(), "post.process")
	defer span.Finish()

	req, err := http.NewRequest("GET", "http://example.com", nil)
	req = req.WithContext(ctx)
	// Inject the span Context in the Request headers
	err = tracer.Inject(span.Context(), tracer.HTTPHeadersCarrier(req.Header))
	if err != nil {
		// Handle or log injection error
	}
	http.DefaultClient.Do(req)

	w.Write([]byte("ok 3"))
}
