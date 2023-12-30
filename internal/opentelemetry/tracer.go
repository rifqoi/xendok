package opentelemetry

import (
	"context"
	"fmt"

	"github.com/rifqoi/xendok-service/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func newTraceProvider(exp sdktrace.SpanExporter) (*sdktrace.TracerProvider, error) {
	// Get config
	serviceName := config.Get().ServiceName
	if serviceName == "" {
		return nil, fmt.Errorf("service_name config is not provided.")
	}

	// Ensure default SDK resources and the required service name are set.
	// r, err := resource.Merge(
	// 	resource.Default(),
	// 	resource.NewWithAttributes(
	// 		semconv.SchemaURL,
	// 		semconv.ServiceName(serviceName),
	// 	),
	// )

	r, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to merge trace resource: %v", err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	), nil
}

func GetSpan(ctx context.Context, status string) trace.Span {
	serviceName := config.Get().ServiceName
	_, span := otel.Tracer(serviceName).Start(ctx, status)

	return span
}
