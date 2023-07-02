package main

import (
	"github.com/SanketPatil6869/Task_Management_Gin1/config"
	"github.com/SanketPatil6869/Task_Management_Gin1/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	/* routes */
	router.POST("/create", controller.CreateTask)
	defer router.Run()
}
