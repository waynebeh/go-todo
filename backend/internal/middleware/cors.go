package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

// CORSConfig returns a configured CORS middleware for Gin
func CORSConfig() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins:     []string{"http://127.0.0.1:5500"}, // Add your frontend's origin here
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    })
}