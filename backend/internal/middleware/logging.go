package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware log detail for each requests
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start the timer
		start := time.Now()

		// Process the request
		c.Next()

		duration := time.Since(start)

		timestamp := time.Now().UTC().Format(time.RFC3339)

		method := c.Request.Method
		path := c.Request.URL.Path
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		queryParams := c.Request.URL.RawQuery
		responseSize := c.Writer.Size()

		log.Printf("timestamp=%s method=%s path=%s status=%d duration=%.3fms client_ip=%s user_agent=%s query_params=%s response_size=%d",
		timestamp,
		method,
		path,
		status,
		float64(duration.Milliseconds()),
		clientIP,
		userAgent,
		queryParams,
		responseSize,
	)

		// show error log 
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				log.Printf("[%s] Level: ERROR, Message: %s, Path: %s, Method: %s", timestamp, err.Error(), path, method)
			}
		}
	}
}
