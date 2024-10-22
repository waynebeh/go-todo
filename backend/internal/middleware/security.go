package middleware

import "github.com/gin-gonic/gin"


// SecurityMiddleware sets security-related headers
func SecurityMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Set security headers
        c.Header("X-Content-Type-Options", "nosniff")
        c.Header("X-Frame-Options", "DENY")
        c.Header("X-XSS-Protection", "1; mode=block")
        c.Header("Content-Security-Policy", "default-src 'self'")

        c.Next()
    }
}