package opentelemetry

import (
	"context"
	"fmt"

	"github.com/rifqoi/xendok-service/internal/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func setDefaultTracer(exporter sdktrace.SpanExporter) error {
	tp, err := newTraceProvider(exporter)
	if err != nil {
		return err
	}
	otel.SetTracerProvider(tp)

	return nil

}

func WithStdoutExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	stdoutExporter, err := stdouttrace.New()
	if err != nil {
		return nil, err
	}

	err = setDefaultTracer(stdoutExporter)
	if err != nil {
		return nil, err
	}

	return stdoutExporter, nil
}

func WithHttpExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	// Get OTLP HTTP Endpoint
	otlpEndpoint := config.Get().Otel.Endpoint.Http
	if otlpEndpoint == "" {
		return nil, fmt.Errorf("otel.endpoint.http config is not provided.")
	}

	// Change default to http
	insecureOpt := otlptracehttp.WithInsecure()
	endpointOpt := otlptracehttp.WithEndpoint(otlpEndpoint)

	httpExporter, err := otlptracehttp.New(ctx, insecureOpt, endpointOpt)
	if err != nil {
		return nil, err
	}

	err = setDefaultTracer(httpExporter)
	if err != nil {
		return nil, err
	}

	return httpExporter, nil
}

func WithGRPCExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	// Get OTLP GRPC Endpoint
	otlpEndpoint := config.Get().Otel.Endpoint.GRPC
	if otlpEndpoint == "" {
		return nil, fmt.Errorf("otel.endpoint.grpc config is not provided.")
	}

	// Change default to http
	insecureOpt := otlptracegrpc.WithInsecure()
	endpointOpt := otlptracegrpc.WithEndpoint(otlpEndpoint)

	grpcExporter, err := otlptracegrpc.New(ctx, insecureOpt, endpointOpt)
	if err != nil {
		return nil, err
	}

	err = setDefaultTracer(grpcExporter)
	if err != nil {
		return nil, err
	}

	return grpcExporter, nil
}
