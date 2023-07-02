package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("task_management.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	DB = database
}
