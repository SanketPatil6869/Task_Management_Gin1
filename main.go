package main

import (
	"github.com/SanketPatil6869/Task_Management_Gin1/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	router.Run()
}
