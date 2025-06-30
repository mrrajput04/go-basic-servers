package config

import (
	"fmt"
	"go-todo-crud/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "test.db" // Default path
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	if err := database.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		panic("Failed to run migration: " + err.Error())
	}
	DB = database
	fmt.Println("Database connection successful and migration complete.")
}
