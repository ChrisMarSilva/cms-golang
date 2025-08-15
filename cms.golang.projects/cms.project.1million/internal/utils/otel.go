package utils

import (
	"context"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	Tracer trace.Tracer
	Meter  metric.Meter
	//PromHandler *metric.MeterProvider
	// MetricHttpRequestCount metric.Int64Counter
	// MetricHttpLatency      metric.Float64Histogram
	// MetricPingRequestCount metric.Int64Counter
	// MetricTotalSalesValue  metric.Float64Counter

	// MetricHttpCounter = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "dental_inventory_access_total",
	// 		Help: "Total number of times inventory items are accessed",
	// 	},
	// 	[]string{"item_type"},
	// )

	// MetricHttpRequestCount = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "myapp_requests_total",
	// 		Help: "Total number of requests processed by the MyApp web server.",
	// 	},
	// 	[]string{"path", "status"},
	// )

	// MetricHttpErrorCount = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "myapp_requests_errors_total",
	// 		Help: "Total number of error requests processed by the MyApp web server.",
	// 	},
	// 	[]string{"path", "status"},
	// )

	MetricHttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	MetricHttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	MetricHttpActiveConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Number of active connections",
		},
	)
)

func InitOpenTelemetry(serviceName string) func() {
	ctx := context.Background()

	// --- Exportador Jaeger ---
	client := otlptracegrpc.NewClient(otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("localhost:4317"))
	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// Recursos compartilhados
	res := resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String(serviceName))

	// Provedor de traces
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp), sdktrace.WithResource(res))

	// --- Exportador Prometheus ---
	metricExp, err := otlpmetricgrpc.New(ctx)
	//metricExp, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}

	// Provedor de métricas
	//mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(metricExp), sdkmetric.WithResource(res))
	mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExp)), sdkmetric.WithResource(res))

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(mp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	Tracer = tp.Tracer("cms.project.1million")
	//Meter = mp.Meter("cms.project.1million")

	// --- Definindo métricas ---
	// MetricHttpRequestCount, _ = Meter.Int64Counter("http_requests_total", metric.WithDescription("Contagem total de requisições HTTP"))
	// MetricHttpLatency, _ = Meter.Float64Histogram("http_request_duration_seconds", metric.WithDescription("Latência das requisições HTTP"))
	// MetricPingRequestCount, _ = Meter.Int64Counter("ping_requests_total", metric.WithDescription("Quantidade de chamadas ao /ping"))
	// MetricTotalSalesValue, _ = Meter.Float64Counter("sales_total_value", metric.WithDescription("Valor total de vendas processadas"))

	//PromHandler = metricExp

	return func() {
		err = tp.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		err = mp.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}
