package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start timer
		t := time.Now()
		path := c.Request.URL.Path

		//process request
		c.Next()

		//log request details
		latency := time.Since(t)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method

		fmt.Printf("[%s] %s | %3d | %13v | %15s | %s\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			method,
			statusCode,
			latency,
			clientIP,
			path,
		)
	}
}
