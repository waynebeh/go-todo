package handler

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Completed   bool      `json:"completed"`
}

var todos = []Todo{
	{ID: 1, Description: "Buy groceries", DueDate: time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC), Completed: false},
	{ID: 2, Description: "Go to the Gym", DueDate: time.Date(2024, time.March, 16, 0, 0, 0, 0, time.UTC), Completed: false},
	{ID: 3, Description: "Eating breakfast", DueDate: time.Date(2024, time.March, 17, 0, 0, 0, 0, time.UTC), Completed: false},
	{ID: 4, Description: "Running", DueDate: time.Date(2024, time.March, 18, 0, 0, 0, 0, time.UTC), Completed: false},
}

// Get all Todos
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// Get a Todo by ID
func GetTodosById(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range todos {
		if fmt.Sprintf("%d", todo.ID) == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// Create a new Todo
func CreateTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add new todo to the slice
	newTodo.ID = len(todos) + 1 // Auto-generate an ID (you can modify this logic if needed)
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

// Update an existing Todo
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find and update the todo
	for index, t := range todos {
		if fmt.Sprintf("%d", t.ID) == id {
			// Update the todo in the slice
			todos[index] = updatedTodo
			todos[index].ID = t.ID // Preserve the original ID
			c.JSON(http.StatusOK, todos[index])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// Delete a Todo by ID
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	for index, t := range todos {
		if fmt.Sprintf("%d", t.ID) == id {
			// Remove the todo from the slice
			todos = append(todos[:index], todos[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}