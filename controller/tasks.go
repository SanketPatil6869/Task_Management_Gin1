package controller

import (
	"net/http"

	"github.com/SanketPatil6869/Task_Management_Gin1/config"
	"github.com/SanketPatil6869/Task_Management_Gin1/models"
	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Due_Date    string `json:"due_date" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

func CreateTask(c *gin.Context) {
	var input CreateTaskInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task := models.Task{Title: input.Title, Description: input.Description, Due_Date: input.Due_Date, Status: input.Status}
	config.DB.Create(&task)
	c.JSON(http.StatusOK, gin.H{"Data": task})
}