package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shixinshuiyou/framework/netx"
	"net/http"
	"strconv"
	"time"
)

// prometheus metrics
// 请求量 http_request_count
// 请求耗时 http_request_duration_seconds
// 请求大小 http_request_size_bytes
// 响应大小 http_response_size_bytes
// panic量 http_response_panic_count
// 还包含 prometheus go 客户端自带的一些指标

// 默认不采集，使用metricMiddleware 时才会采集

var RecoveryMetric = func(c *gin.Context, err interface{}) {}

// MetricMiddleware 采集prometheus日志
func MetricMiddleware(prefix string) gin.HandlerFunc {
	ns := "service"
	if prefix != "" {
		ns = prefix + "_" + ns
	}

	labels := []string{"idc", "endpoint", "status", "method", "path"}
	reqCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: ns,
		Name:      "http_request_count",
		Help:      "Total number of server requests.",
	}, labels)

	reqDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: ns,
		Name:      "http_request_duration_seconds",
		Help:      "HTTP request latencies in seconds",
	}, labels)

	reqSizeBytes := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: ns,
		Name:      "http_request_size_bytes",
		Help:      "HTTP request sizes in bytes.",
	}, labels)

	respSizeBytes := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: ns,
		Name:      "http_response_size_bytes",
		Help:      "HTTP response sizes in bytes",
	}, labels)

	panicCount := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: ns,
		Name:      "http_response_panic_count",
		Help:      "Total number of panic response",
	}, labels)
	RecoveryMetric = func(ctx *gin.Context, err interface{}) {
		lvs := []string{string(netx.IDC()), netx.InternalIp(), strconv.Itoa(ctx.Writer.Status()), ctx.Request.Method, ctx.Request.URL.Path}
		panicCount.WithLabelValues(lvs...).Inc()
	}

	prometheus.MustRegister(reqCount, reqDuration, reqSizeBytes, respSizeBytes, panicCount)

	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		lvs := []string{string(netx.IDC()), netx.InternalIp(), strconv.Itoa(ctx.Writer.Status()), ctx.Request.Method, ctx.Request.URL.Path}

		reqCount.WithLabelValues(lvs...).Inc()
		reqDuration.WithLabelValues(lvs...).Observe(time.Since(start).Seconds())
		reqSizeBytes.WithLabelValues(lvs...).Observe(calcRequestSize(ctx.Request))
		respSizeBytes.WithLabelValues(lvs...).Observe(float64(ctx.Writer.Size()))
	}
}

func calcRequestSize(request *http.Request) float64 {
	size := 0
	if request.URL != nil {
		size = len(request.URL.String())
	}

	size += len(request.Method)
	size += len(request.Proto)

	for name, values := range request.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(request.Host)

	// request.Form and request.MultipartForm are assumed to be included in request.URL.
	if request.ContentLength != -1 {
		size += int(request.ContentLength)
	}
	return float64(size)
}
