package trace

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.12.0"
	"strconv"
)

var EmptyConfig Config

type Config struct {
	Host        string
	Port        int
	ServiceName string
	Ratio       float64
}

func NewJaegerTracer(config Config) (*trace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost(config.Host), jaeger.WithAgentPort(strconv.Itoa(config.Port))))
	if err != nil {
		return nil, err
	}
	provider := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(config.Ratio)),
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
		)),
	)
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return provider, nil
}
