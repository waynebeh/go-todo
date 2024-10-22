package main

import (
	"Todo-go-htmx/internal/handler"
	"Todo-go-htmx/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// middleware
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.CORSConfig())
	r.Use(middleware.SecurityMiddleware())

	// routes
	r.GET("/todos", handler.GetTodos)
	r.GET("/todos/:id", handler.GetTodosById)
	r.POST("/todos", handler.CreateTodo)     
	r.PUT("/todos/:id", handler.UpdateTodo)  
	r.DELETE("/todos/:id", handler.DeleteTodo)

	// run on localhost:5000
	r.Run(":5000")
}
