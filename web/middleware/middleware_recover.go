package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
)

func RecoverMiddleware(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			RecoveryMetric(c, rec)
			logrus.WithContext(c).Error(c.Request.URL)
			err := fmt.Sprintf("%v", rec)
			logrus.WithContext(c).Errorf("err: %s, stack: %s", err, string(debug.Stack()))
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": err,
			})
			return
		}
		//Recover
	}(c)
	c.Next()
}
