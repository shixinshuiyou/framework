package tracer

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const contextTracerKey = "Tracer-context"

// Jaeger 通过 middleware 将 tracer 和 ctx 注入到 gin.Context 中
func Jaeger(srvName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var sp opentracing.Span
		var md = make(map[string]string)
		tracer := opentracing.GlobalTracer()
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err == nil {
			sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
			tracer = sp.Tracer()
		}
		defer sp.Finish()


		tracer.Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(md))

		ctx := context.TODO()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		c.Set(contextTracerKey, ctx)

		c.Next()

		ext.HTTPStatusCode.Set(sp, uint16(c.Writer.Status()))
		ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
		ext.HTTPMethod.Set(sp, c.Request.Method)
		// 把tracer 写入到context中
	}
}

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {
	if v, exist := c.Get(contextTracerKey); exist {
		ctx, ok = v.(context.Context)
		return
	}
	return context.TODO(), false
}
