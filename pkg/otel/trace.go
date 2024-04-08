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

type OTel struct {
	traceP *sdktrace.TracerProvider
	tracer trace.Tracer

	shutdown func(ctx context.Context) error
}

type Config struct {
	Name, Version string
}

func MustNew(ctx context.Context, cfg *Config) *OTel {
	otl, err := New(ctx, cfg)
	if err != nil {
		panic(err)
	}
	return otl
}

func New(ctx context.Context, cfg *Config) (*OTel, error) {
	var otl OTel
	if cfg.Name == "" {
		nop := noop.NewTracerProvider()
		otl.tracer = nop.Tracer("no-op-provider")
		return &otl, nil
	}

	tp, err := newTraceProvider(ctx, cfg)
	otl.traceP = tp
	otl.shutdown = tp.Shutdown
	if err != nil {
		err = errors.Join(err, otl.shutdown(ctx))
		return &otl, err
	}
	otel.SetTextMapPropagator(
		newPropagator(),
	)
	otel.SetTracerProvider(otl.traceP)
	otl.tracer = otl.traceP.Tracer(
		cfg.Name,
		trace.WithInstrumentationVersion(cfg.Version),
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

func newTraceProvider(ctx context.Context, cfg *Config) (*sdktrace.TracerProvider, error) {
	exporter, err := newExporter(ctx)
	if err != nil {
		return nil, fmt.Errorf("new exporter: %w", err)
	}
	traceP := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter,
			sdktrace.WithBatchTimeout(time.Second)),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Name),
			semconv.ServiceVersionKey.String(cfg.Version),
		)),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(100)),
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
