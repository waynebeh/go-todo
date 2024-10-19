package main

import (
	"Todo-go-htmx/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.Static("/static","/web/static")

	// r.GET("/", func(c * gin.Context){
	// 	c.File("./web/static/index.html")
	// })
	r.GET("/todos", handler.GetTodos)
	r.GET("/todos/:id", handler.GetTodosById)
	r.POST("/todos", handler.CreateTodo)     
	r.PUT("/todos/:id", handler.UpdateTodo)  
	r.DELETE("/todos/:id", handler.DeleteTodo)

	// run on localhost:5000
	r.Run(":5000")
}
