package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// go mod init github.com/chrismarsilva/cms.opentelemetry
// go get -u "go.opentelemetry.io/otel"
// go get -u "go.opentelemetry.io/otel/sdk/resource"
// go get -u "go.opentelemetry.io/otel/sdk/trace"
// go get -u "go.opentelemetry.io/otel/sdk/metric"
// go get -u "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
// go get -u "go.opentelemetry.io/otel/exporters/otlp"
// go get -u "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
// go get -u "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
// go get -u "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
// go get -u "go.opentelemetry.io/otel/exporters/jaeger"
// go get -u "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
// go get -u "go.opentelemetry.io/otel/propagation"
// go get -u "go.opentelemetry.io/otel/attribute"
// go get -u "go.opentelemetry.io/otel/codes"
// go get -u "go.opentelemetry.io/otel/metric"
// go get -u "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
// go get -u "google.golang.org/grpc"
// go get -u "google.golang.org/grpc/credentials/insecure"
// go get -u "github.com/chrismarsilva/cms.opentelemetry/bridges/otelslog"
// go mod tidy

// docker-compose down
// docker-compose up -d --build

// k6 run test.js

// go run main.go

var (
	name = "Example-Go-Tracer" // "go.opentelemetry.io/contrib/examples/otel-collector" // "Example-Go-Tracer"
	tr   = otel.Tracer(name)
	//mt          = otel.Meter(name)
	//viewCounter metric.Int64Counter
)

// func init() {
// 	var err error
// 	viewCounter, err = mt.Int64Counter("user.views", metric.WithDescription("The number of views"), metric.WithUnit("{views}"))
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	ctx := context.Background()
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	cleanup := initTracer(ctx)
	defer cleanup()

	// ctx, span := tr.Start(ctx, "main")
	// defer span.End()
	// doWork(ctx)

	router := http.NewServeMux()
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/pay", payHandler)
	router.Handle("/start", otelhttp.NewHandler(http.HandlerFunc(startHandler), "CheckoutStart"))
	router.Handle("/finish", otelhttp.NewHandler(http.HandlerFunc(finishHandler), "CheckoutFinish"))
	router.Handle("/hello", otelhttp.NewHandler(http.HandlerFunc(helloHandler), "GET /hello"))
	//router.Handle("/metrics", promhttp.Handler())
	//router.HandleFunc("/metrics", payHandler)

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		router.Handle(pattern, handler)
	}
	handleFunc("/rolldice/", rolldice)

	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      otelhttp.NewHandler(router, "/"),
	}

	//log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(srv.ListenAndServe())
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

func rolldice(w http.ResponseWriter, r *http.Request) {
	_, span := tr.Start(r.Context(), "rolldiceeee")
	defer span.End()

	span.SetAttributes(attribute.String("key", "value"))
	span.AddEvent("eventName", trace.WithAttributes(attribute.String("key", "value")))

	//viewCounter.Add(ctx, 1)

	// get the current span by the request context
	currentSpan := trace.SpanFromContext(r.Context())
	currentSpan.SetAttributes(attribute.String("hello.name", name))

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("rolldice"))
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
	// "otel-collector:4317" "localhost:4317"
	conn, err := grpc.NewClient("localhost:4317", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC connection to collector: %v", err)
	}

	// Sets up OTLP HTTP exporter with endpoint, headers, and TLS config.
	//exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	//exp, err := otlptrace.New(ctx, otlptracehttp.NewClient(otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithHeaders(map[string]string{"content-type": "application/json"}), otlptracehttp.WithInsecure()))
	//exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn), otlptracegrpc.WithInsecure())
	//exp, err := otlptracegrpc.New(ctx)
	if err != nil {
		log.Fatalf("failed to create trace exporter: %v", err)
	}

	// Defines resource with service name, version, and environment.
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("CMS-Example-Trace"))
	if err != nil {
		log.Fatalf("Failed to create resource: %v", err)
	}

	// Configures the tracer provider with the exporter and resource.
	// sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithResource(res), sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exp)),	)
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp), sdktrace.WithResource(res))
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// metricExp, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	// if err != nil {
	// 	log.Fatalf("failed to create metrics exporter: %v", err)
	// }

	// mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExp)), sdkmetric.WithResource(res))
	// otel.SetMeterProvider(mp)

	return func() {
		_ = tp.Shutdown(context.Background())
		//_ = mp.Shutdown(context.Background())
	}
}
