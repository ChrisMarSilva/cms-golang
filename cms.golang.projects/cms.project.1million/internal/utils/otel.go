package utils

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

const Name = "github.com/chrismarsilva/cms.project.1million"

var (
	Tracer = otel.Tracer(Name)
	// Meter  otel.Meter(Name)
	Logger = otelslog.NewLogger(Name)
	// Logger *slog.Logger // = otelslog.NewLogger(Name)

	// MetricHttpRequestsTotal = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "http_requests_total",
	// 		Help: "Total number of HTTP requests",
	// 	},
	// 	[]string{"path"}, // 	[]string{"path", "status"}, // 	[]string{"item_type"},
	// )

	// MetricHttpRequestDuration = prometheus.NewHistogramVec(
	// 	prometheus.HistogramOpts{
	// 		Name:    "http_request_duration_seconds",
	// 		Help:    "Duration of HTTP requests",
	// 		Buckets: prometheus.DefBuckets,
	// 	},
	// 	[]string{"path"},
	// )

	// MetricHttpActiveConnections = prometheus.NewGauge(
	// 	prometheus.GaugeOpts{
	// 		Name: "active_connections",
	// 		Help: "Number of active connections",
	// 	},
	// )
)

func InitOpenTelemetry(serviceName string) func() {
	ctx := context.Background()

	// Configura propagadores

	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)

	otel.SetTextMapPropagator(prop)

	// Recursos do serviço

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String("1.0.0"))

	// client

	// conn, err := grpc.NewClient(
	// 	"localhost:4317",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	slog.Error("Failed to create gRPC client", slog.Any("error", err))
	// 	os.Exit(1)
	// }

	// --- Exportador Jaeger ---

	// traceExporter, err := otlptracegrpc.New(
	// 	ctx,
	// 	otlptracegrpc.WithGRPCConn(conn),
	// )

	traceExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"), // OTLP/gRPC
		// otlptracegrpc.WithEndpoint("http://localhost:4318"),  // OTLP/HTTP
	)
	if err != nil {
		slog.Error("Failed to create trace exporter", slog.Any("error", err))
		os.Exit(1)
	}

	// Provedor de traces

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res))

	otel.SetTracerProvider(tp)

	// Tracer = tp.Tracer("cms.project.1million") // otel.Tracer

	// --- Exportador Prometheus ---

	// metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	// if err != nil {
	// 	slog.Error("Failed to create Prometheus exporter", slog.Any("error", err))
	// 	os.Exit(1)
	// }

	// Provedor de métricas

	// mp := sdkmetric.NewMeterProvider(
	// 	//sdkmetric.WithReader(metricExporter)
	// 	sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
	// 	sdkmetric.WithResource(res))

	// otel.SetMeterProvider(mp)

	// Meter = mp.Meter("cms.project.1million") // otel.Meter

	// --- Exportador logger ---

	// logExporter, err := otlploggrpc.New(
	// 	ctx,
	// 	otlploggrpc.WithInsecure(),
	// 	otlploggrpc.WithEndpoint("localhost:4317"), // OTLP/gRPC
	// 	// otlploggrpc.WithEndpoint("http://localhost:4318"),  // OTLP/HTTP
	// )

	// logExporter, err := otlploggrpc.New(
	// 	ctx,
	// 	otlploggrpc.WithGRPCConn(conn),
	// )

	logExporter, err := otlploggrpc.New(
		ctx,
		otlploggrpc.WithInsecure(),
		otlploggrpc.WithEndpoint("localhost:4317"), // Collector OTLP/gRPC
		// otlploggrpc.WithEndpoint("http://localhost:4318"),  // Collector OTLP/HTTP
	)

	if err != nil {
		slog.Error("Failed to create OTLP logger exporter", slog.Any("error", err))
		os.Exit(1)
	}

	// Provedor de logger

	lp := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
		log.WithResource(res),
	)

	global.SetLoggerProvider(lp)

	Logger = otelslog.NewLogger(
		Name,
		otelslog.WithSource(true),
		otelslog.WithLoggerProvider(lp),
	)

	return func() {
		err = tp.Shutdown(context.Background())
		if err != nil {
			slog.Error("Failed to shutdown tracer provider", slog.Any("error", err))
		}

		// err = mp.Shutdown(context.Background())
		// if err != nil {
		// 	slog.Error("Failed to shutdown meter provider", slog.Any("error", err))
		// }

		err = lp.Shutdown(context.Background())
		if err != nil {
			slog.Error("Failed to shutdown logger provider", slog.Any("error", err))
		}
	}
}
