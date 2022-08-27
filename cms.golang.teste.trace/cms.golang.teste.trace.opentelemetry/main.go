package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.trace.opentelemetry
// go get go.opentelemetry.io/otel
// go get go.opentelemetry.io/otel/trace
// go get go.opentelemetry.io/otel/sdk
// go get go.opentelemetry.io/otel/exporters/jaeger
// go get go.opentelemetry.io/otel/exporters/otlp/otlptrace
// go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
// go get github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace
// go mod tidy

// go run main.go

const (
	service     = "cms.teste.trace"
	environment = "production"
	id          = 1
)

func main() {
	log.Println("Main.Ini")

	ctx := context.Background()

	//tp, err := tracerProvider("http://localhost:14268/api/traces")
	tp, err := tracerProvider("http://localhost:16686/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	tr := tp.Tracer("component-main")

	ctx, span := tr.Start(ctx, "foo")
	defer span.End()

	bar(ctx)

	newCtx, span2 := otel.Tracer("component-1").Start(ctx, "Run")
	defer span2.End()

	_, spanP := otel.Tracer("component-2").Start(newCtx, "Poll")
	defer spanP.End()

	_, spanW := otel.Tracer("component-3").Start(newCtx, "Write")
	defer spanW.End()

	log.Println("Main.Fim")
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {

	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)

	return tp, nil
}

func bar(ctx context.Context) {
	// Use the global TracerProvider.
	tr := otel.Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	span.SetAttributes(attribute.Key("testset").String("value"))
	defer span.End()

	// Do bar...
}
