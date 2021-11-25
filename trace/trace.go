package trace

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

const JaegerHostPort = "127.0.0.1:6831"

// initJaegerTracer returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func initJaegerTracer(serviceName string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			//LocalAgentHostPort: agentHostPort,
		},
	}

	sender, _ := jaeger.NewUDPTransport(JaegerHostPort, 0)

	reporter := jaeger.NewRemoteReporter(sender)
	// Initialize tracer with a logger and a metrics factory
	tracer, closer, _ := cfg.NewTracer(
		config.Reporter(reporter),
		config.Logger(log.StdLogger),
		config.Metrics(metrics.NullFactory),
		// 设置最大 Tag 长度，根据情况设置
		config.MaxTagValueLength(65535),
	)
	return tracer, closer
}

//
func OpenGinTracing(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ginTracing, closer := initJaegerTracer(serviceName)
		defer closer.Close()
		spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		span := opentracing.StartSpan(
			c.Request.URL.Path,
			opentracing.ChildOf(spanCtx),
			opentracing.Tag{Key: string(ext.Component), Value: "Http"},
			ext.SpanKindRPCServer,
		)
		defer span.Finish()
		c.Set("tracer-context", ginTracing)
		c.Set("span-context", span.Context())
		c.Next()
	}
}
