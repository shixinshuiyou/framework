package middleware

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	_const "github.com/shixinshuiyou/framework/const"
	"github.com/shixinshuiyou/framework/util"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"time"
)

// CustomResponseWriter 封装 gin ResponseWriter 用于获取回包内容。
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func MiddlewareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(_const.TraceID)
		if traceID == "" {
			// 如果为空则创建
			traceID = strconv.FormatInt(util.GetSnowflakeID(), 10)
		}
		// 将request_id保存
		c.Set(_const.TraceID, traceID)
		ctx := c.Request.Context()
		context.WithValue(ctx, _const.TraceID, traceID)
		c.Request = c.Request.WithContext(ctx)
		// 使用自定义 ResponseWriter
		start := time.Now()
		reqBody, _ := c.GetRawData()
		// 解决:使用该函数回导致参数被读取，后续使用Bind、BIndJson会出现EOF问题
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		logrus.WithFields(logrus.Fields{
			"trace_id": traceID,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"req":      string(reqBody),
		}).Info()

		// 使用自定义 ResponseWriter
		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = crw

		c.Next()

		logrus.WithFields(logrus.Fields{
			"trace_id": traceID,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"resp":     string(crw.body.Bytes()),
			"latency":  time.Now().Sub(start).Milliseconds(),
		}).Debug()

		// 将request_id写到resp中
		c.Header(_const.TraceID, traceID)
	}
}
