package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/framework/netx"
	"os"
)

// SvcTagMiddleware IP+进程名
func SvcTagMiddleware(c *gin.Context) {
	ip := netx.InternalIp()
	processID := os.Getpid() // 进程ID
	c.Writer.Header().Add("PLAT-X-SER", fmt.Sprintf("%s-%d", ip, processID))
}
