package trace

import (
	"github.com/sirupsen/logrus"
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

func InitJaegerTracer(config Config) (*trace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost(config.Host), jaeger.WithAgentPort(strconv.Itoa(config.Port))))
	if err != nil {
		return nil, err
	}
	logrus.Debugf("init jaeger %s:%d", config.Host, config.Port)
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
