package handlers

import (
	"tms-backend/src/database"
	"tms-backend/src/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	userID := c.MustGet("id").(uint)

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = userID
	database.DB.Create(&todo)
	c.JSON(201, todo)
}

func GetTodos(c *gin.Context) {
	userID := c.MustGet("id").(uint)
	var todos []models.Todo
	database.DB.Where("user_id = ?", userID).Find(&todos)
	c.JSON(200, todos)
}
