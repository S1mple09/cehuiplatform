package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求开始时记录时间
		startTime := time.Now()

		// 继续处理请求
		c.Next()

		// 记录请求的详细信息
		_ = time.Since(startTime)

	}
}
