// In a hypothetical main.go
package main

import (
	"go-todo-crud/config"
	"go-todo-crud/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create handler with dependencies
	h := &handlers.Handler{DB: db}

	router := gin.Default()

	// Register routes with handler methods
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
	// ... other routes for Todos

	router.Run("localhost:8080")
}
