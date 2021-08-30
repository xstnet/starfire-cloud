package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestCostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("requestStartTime", time.Now())
		//请求处理
		c.Next()
		// url := c.Request.URL.String()
	}
}
