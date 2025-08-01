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
	//exp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
	//exp, err := otlptrace.New(ctx, otlptracegrpc.NewClient(otlptracegrpc.WithEndpoint("localhost:4318"), otlptracegrpc.WithInsecure()))
	//exp, err := otlptrace.New(ctx, otlptracehttp.NewClient(otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure()))
	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn), otlptracegrpc.WithInsecure())
	//exp, err := otlptracegrpc.New(ctx)
	if err != nil {
		log.Fatalf("failed to create trace exporter: %v", err)
	}

	// Defines resource with service name, version, and environment.
	// res, err := resource.New(ctx, resource.WithAttributes(attribute.String("service.name", "CMS-Example-Trace")))
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("CMS-Example-Trace"))
	if err != nil {
		log.Fatalf("Failed to create resource: %v", err)
	}

	// Configures the tracer provider with the exporter and resource.
	// sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithResource(res), sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exp)),	)
	//tp := sdktrace.NewTracerProvider(sdktrace.WithSampler(sdktrace.AlwaysSample()), sdktrace.WithBatcher(exp), sdktrace.WithResource(res))
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

/*


provider := metrics.InitMeter()
	meter := provider.Meter("sample-golang-app")
	metrics.GenerateMetrics(meter)

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName))

func FindBooks(c *gin.Context) {
	var books []models.Book
	span := trace.SpanFromContext(c.Request.Context())
	span.SetAttributes(attribute.String("controller", "books"))
	span.AddEvent("This is a sample event", trace.WithAttributes(attribute.Int("pid", 4328), attribute.String("sampleAttribute", "Test")))


package metrics

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	api "go.opentelemetry.io/otel/metric"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc/credentials"
)

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

// Example use cases for sync counter
// - count the number of bytes received
// - count the number of requests completed
// - count the number of accounts created
// - count the number of checkpoints run
// - count the number of HTTP 5xx errors
//
// The increments should be non-negative.

func exceptionsCounter(meter api.Meter) {
	counter, err := meter.Int64Counter("exceptions", api.WithUnit("1"),
		api.WithDescription("Counts exceptions since start"),
	)
	if err != nil {
		log.Fatal("Error in creating exceptions counter: ", err)
	}

	for {
		// Increment the counter by 1.
		// The attributes describe the exception.
		counter.Add(context.Background(), 1, api.WithAttributes(attribute.KeyValue{
			Key: attribute.Key("exception_type"), Value: attribute.StringValue("NullPointerException"),
		}))
		time.Sleep(time.Duration(rand.Int63n(5)) * time.Millisecond)
	}
}

// Example use cases for async counter
//   - count the number of page faults
//   - CPU time, which could be reported for each thread, each process or the
//     entire system. For example "the CPU time for process
//     A running in user mode, measured in seconds".
//
// Basically, any value that is monotonically increasing and happens in the background.
// The increments should be non-negative.
func pageFaultsCounter(meter api.Meter) {
	counter, err := meter.Int64ObservableCounter(
		"page_faults",
		api.WithUnit("1"),
		api.WithDescription("Counts page faults since start"),
	)
	if err != nil {
		log.Fatal("Error in creating page faults counter: ", err)
	}

	_, err = meter.RegisterCallback(
		func(_ context.Context, o api.Observer) error {
			attrSet := attribute.NewSet(attribute.String("process", "boo"))
			withAttrSet := api.WithAttributeSet(attrSet)
			o.ObserveInt64(counter, rand.Int63n(100), withAttrSet)
			return nil
		},
		counter,
	)
	if err != nil {
		log.Fatal("Error in registering callback: ", err)
	}
}

// Example use cases for Histogram
// - the request duration
// - the size of the response payload
func requestDurationHistogram(meter api.Meter) {
	histogram, err := meter.Int64Histogram(
		"http_request_duration",
		api.WithUnit("ms"),
		api.WithDescription("The HTTP request duration in milliseconds"),
	)
	if err != nil {
		log.Fatal("Error in creating request duration histogram: ", err)
	}

	for {
		histogram.Record(context.Background(), rand.Int63n(1000), api.WithAttributes(attribute.String("path", "/api/boo")))
		time.Sleep(time.Duration(rand.Int63n(5)) * time.Millisecond)
	}
}

// Asynchronous Gauge is an Instrument
// which reports non-additive value(s)
// (e.g. the room temperature - it makes no sense to report the
// temperature value from multiple rooms and sum them up) when the
// instrument is being observed.

// Example use cases for Async Gauge
// - the current room temperature
// - the CPU fan speed

// Note: if the values are additive (e.g. the process heap size -
// it makes sense to report the heap size from multiple processes and sum them up,
// so we get the total heap usage),
// use Asynchronous Counter or Asynchronous UpDownCounter.

func roomTemperatureGauge(meter api.Meter) {
	gauge, err := meter.Float64ObservableGauge(
		"room_temperature",
		api.WithUnit("1"),
		api.WithDescription("The room temperature in celsius"),
	)
	if err != nil {
		log.Fatal("Error in creating room temperature gauge: ", err)
	}

	_, err = meter.RegisterCallback(
		func(_ context.Context, o api.Observer) error {
			attrSet := attribute.NewSet(attribute.String("process", "boo"))
			withAttrSet := api.WithAttributeSet(attrSet)
			o.ObserveFloat64(gauge, rand.Float64()*100, withAttrSet)
			return nil
		},
		gauge,
	)
	if err != nil {
		log.Fatal("Error in registering callback: ", err)
	}
}

// UpDownCounter is an Instrument which supports increments and decrements.
// if the value is monotonically increasing, use Counter instead.
// Example use cases for UpDownCounter
// - the number of active requests
// - the number of items in a queue

func itemsInQueueUpDownCounter(meter api.Meter) {
	counter, err := meter.Int64UpDownCounter(
		"items_in_queue",
		api.WithUnit("1"),
		api.WithDescription("The number of items in the queue"),
	)
	if err != nil {
		log.Fatal("Error in creating items in queue up down counter: ", err)
	}

	for {
		counter.Add(context.Background(), rand.Int63n(100), api.WithAttributes(attribute.String("queue", "A")))
		time.Sleep(time.Duration(rand.Int63n(5)) * time.Millisecond)
	}
}

// Asynchronous UpDownCounter is an asynchronous Instrument
// which reports additive value(s)
// (e.g. the process heap size - it makes sense to report the heap size
// from multiple processes and sum them up, so we get the total heap usage)
// when the instrument is being observed.
//
// Example use cases for Asynchronous UpDownCounter
// - the process heap size
// - the approximate number of items in a lock-free circular buffer

func processHeapSizeUpDownCounter(meter api.Meter) {
	counter, err := meter.Float64ObservableUpDownCounter(
		"process_heap_size",
		api.WithUnit("1"),
		api.WithDescription("The process heap size"),
	)
	if err != nil {
		log.Fatal("Error in creating process heap size up down counter: ", err)
	}

	_, err = meter.RegisterCallback(
		func(_ context.Context, o api.Observer) error {
			attrSet := attribute.NewSet(attribute.String("process", "boo"))
			withAttrSet := api.WithAttributeSet(attrSet)
			o.ObserveFloat64(counter, rand.Float64()*100, withAttrSet)
			return nil
		},
		counter,
	)
	if err != nil {
		log.Fatal("Error in registering callback: ", err)
	}
}

func InitMeter() *metricsdk.MeterProvider {

	secureOption := otlpmetricgrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlpmetricgrpc.WithInsecure()
	}

	exporter, err := otlpmetricgrpc.New(
		context.Background(),
		secureOption,
		otlpmetricgrpc.WithEndpoint(collectorURL),
	)

	if err != nil {
		log.Fatalf("Failed to create exporter: %v", err)
	}

	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Fatalf("Could not set resources: %v", err)
	}

	// Register the exporter with an SDK via a periodic reader.
	provider := metricsdk.NewMeterProvider(
		metricsdk.WithResource(res),
		metricsdk.WithReader(metricsdk.NewPeriodicReader(exporter)),
	)
	return provider
}

func GenerateMetrics(meter api.Meter) {
	go exceptionsCounter(meter)
	go pageFaultsCounter(meter)
	go requestDurationHistogram(meter)
	go roomTemperatureGauge(meter)
	go itemsInQueueUpDownCounter(meter)
	go processHeapSizeUpDownCounter(meter)
}
*/
