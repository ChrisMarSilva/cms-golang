package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

// go mod init github.com/chrismarsilva/cms.opentelemetry
// go get -u "go.opentelemetry.io/otel"
// go get -u "go.opentelemetry.io/otel/sdk/resource"
// go get -u "go.opentelemetry.io/otel/sdk/trace"
// go get -u "go.opentelemetry.io/otel/sdk/metric"
// go get -u "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
// go get -u "go.opentelemetry.io/otel/exporters/otlp"
// go get -u "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
// go get -u "go.opentelemetry.io/otel/exporters/jaeger"
// go get -u "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
// go get -u "go.opentelemetry.io/otel/propagation"
// go get -u "go.opentelemetry.io/otel/attribute"
// go get -u "go.opentelemetry.io/otel/codes"
// go get -u "go.opentelemetry.io/otel/metric"
// go get -u "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
// go get -u "google.golang.org/grpc"
// go get -u "google.golang.org/grpc/credentials/insecure"
// go mod tidy

// docker-compose down
// docker-compose up -d --build

// k6 run test.js

// go run main.go

var (
	name = "go.opentelemetry.io/contrib/examples/otel-collector" // "Example-Go-Tracer"
	tr   = otel.Tracer(name)
	mt   = otel.Meter(name)
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cleanup := initTracer(ctx)
	defer cleanup()

	ctx, span := tr.Start(ctx, "main")
	defer span.End()

	doWork(ctx)

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/metrics", payHandler)
	http.HandleFunc("/pay", payHandler)
	http.Handle("/start", otelhttp.NewHandler(http.HandlerFunc(startHandler), "CheckoutStart"))
	http.Handle("/finish", otelhttp.NewHandler(http.HandlerFunc(finishHandler), "CheckoutFinish"))

	//http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tr.Start(r.Context(), "hello")
	defer span.End()

	time.Sleep(500 * time.Millisecond)

	w.Write([]byte("Hello, World!"))
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := tr.Start(r.Context(), "PaymentsService")
	defer span.End()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

	span.AddEvent("Calling Payments Service")

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	res, err := client.Do(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return
	}

	w.Write([]byte(string(body)))
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("checkout started"))
}

func finishHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("checkout finished"))
}

func doWork(ctx context.Context) {
	ctx, span := tr.Start(ctx, "doWork")
	defer span.End()

	time.Sleep(100 * time.Millisecond)

	doSubWork(ctx)
}

func doSubWork(ctx context.Context) {
	ctx, span := tr.Start(ctx, "doSubWork")
	defer span.End()

	span.AddEvent("Performing repository operation")
	time.Sleep(500 * time.Millisecond)
	span.AddEvent("Sub operation completed")

	doSubSubWork(ctx)
}

func doSubSubWork(ctx context.Context) {
	ctx, span := tr.Start(ctx, "doSubSubWork")
	defer span.End()

	span.AddEvent("Performing sub-sub operation")
	time.Sleep(1 * time.Second)
	span.AddEvent("Sub-sub operation completed")
}

func initTracer(ctx context.Context) func() {

	conn, err := grpc.NewClient("localhost:4317", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC connection to collector: %w", err)
	}

	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("CMS-Example-Trace"))
	//res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String("CMS-Example-Trace")))
	// if err != nil {
	// 	log.Fatalf("Failed to create resource: %v", err)
	// }

	//exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
	//exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	//exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		log.Fatalf("failed to create trace exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exp)),
	)

	// tp := sdktrace.NewTracerProvider(
	// 	sdktrace.WithBatcher(exp),
	// 	sdktrace.WithResource(res),
	// )

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		log.Fatalf("failed to create metrics exporter: %v", err)
	}

	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(mp)

	return func() {
		tp.Shutdown(context.Background())
		mp.Shutdown(context.Background())
	}
}
