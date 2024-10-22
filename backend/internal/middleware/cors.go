package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSConfig returns a configured CORS middleware for Gin
func CORSConfig() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins:     []string{"http://127.0.0.1:5500"}, // Frontend origin
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    })
}