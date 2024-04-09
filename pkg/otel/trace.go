package otel

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
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

var globalTracer = &atomic.Value{}

func Tracer() trace.Tracer {
	v, ok := globalTracer.Load().(trace.Tracer)
	if !ok {
		fmt.Println("unable to cast tracer")
		return nil
	}
	return v
}

type Config struct {
	Name, Version string
}

func New(ctx context.Context, cfg *Config) (func(ctx context.Context) error, error) {
	if cfg.Name == "" {
		nop := noop.NewTracerProvider().Tracer("no-op")
		globalTracer.Store(nop)
		return nil, nil
	}

	tp, err := newTraceProvider(ctx, cfg)
	if err != nil {
		err = errors.Join(err, tp.Shutdown(ctx))
		return tp.Shutdown, err
	}
	otel.SetTextMapPropagator(
		newPropagator(),
	)
	otel.SetTracerProvider(tp)
	tracer := otel.Tracer(
		cfg.Name,
		trace.WithInstrumentationVersion(cfg.Version),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
	globalTracer.Store(tracer)
	return tp.Shutdown, nil
}

func newTraceProvider(ctx context.Context, cfg *Config) (*sdktrace.TracerProvider, error) {
	exporter, err := otlptrace.New(ctx, otlptracegrpc.NewClient())
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
