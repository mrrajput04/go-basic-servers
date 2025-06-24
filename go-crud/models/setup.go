package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	DB = db
	fmt.Println("Successfully connected to SQLite database")
}
