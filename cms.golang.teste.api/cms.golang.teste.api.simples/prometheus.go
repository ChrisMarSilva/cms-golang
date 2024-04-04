package main

// import (
// 	"context"

// 	"go.opentelemetry.io/otel/attribute"
// 	"go.opentelemetry.io/otel/sdk/resource"
// )

// func Bootstrap() error {

// 	// Set up trace resource.
// 	res, err := newResource("myApp", "1.0.0")
// 	if err != nil {
// 		return err
// 	}

// 	// Set up trace provider.
// 	tracerProvider, err := newTraceProvider()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func newResource(service string, version string) (*resource.Resource, error) {
// 	return resource.New(
// 		context.Background(),
// 		resource.WithAttributes(attribute.String("service.name", service)),
// 		resource.WithAttributes(attribute.String("service.version", version)),
// 	)
// }
