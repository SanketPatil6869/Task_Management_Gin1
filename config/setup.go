package config

import (
	"github.com/SanketPatil6869/Task_Management_Gin1/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

//var DB2 *sql.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("task_management.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	database.AutoMigrate(&models.Task{})
	DB = database
}
