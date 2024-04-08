package otel

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

const (
	name    = "github.com/eyazici90/go-ddd"
	version = "1.0.0"
)

type OTel struct {
	traceP trace.TracerProvider
	tracer trace.Tracer

	shutdown func(ctx context.Context) error
}

type Config struct {
	SvcName string
}

func New(ctx context.Context, cfg *Config) (*OTel, error) {
	var otl OTel
	if cfg.SvcName == "" {
		nop := noop.NewTracerProvider()
		otl.tracer = nop.Tracer("no-op-provider")
		return &otl, nil
	}

	traceP, err := newTraceProvider(ctx)
	otl.shutdown = traceP.Shutdown
	if err != nil {
		err = errors.Join(err, otl.shutdown(ctx))
		return &otl, err
	}
	otel.SetTextMapPropagator(
		newPropagator(),
	)
	otel.SetTracerProvider(traceP)
	otl.tracer = otl.traceP.Tracer(
		name,
		trace.WithInstrumentationVersion(version),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
	return &otl, nil
}

func (ot *OTel) Tracer() trace.Tracer {
	return ot.tracer
}

func (ot *OTel) Shutdown() func(ctx context.Context) error {
	return ot.shutdown
}

func newTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exporter, err := newExporter(ctx)
	if err != nil {
		return nil, fmt.Errorf("new exporter: %w", err)
	}
	traceP := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter,
			sdktrace.WithBatchTimeout(time.Second)),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
			semconv.ServiceVersionKey.String(version),
		)),
	)
	return traceP, nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	client := otlptracegrpc.NewClient()
	exporter, err := otlptrace.New(ctx, client)
	return exporter, err
}
