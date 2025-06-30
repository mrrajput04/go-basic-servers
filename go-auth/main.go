package main

import (
	"go-auth/handlers"
	"go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Public routes (do not require authentication)
	publicRoutes := r.Group("/public")
	{
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/register", handlers.Register)
	}

	// Protected routes (require authentication)
	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		//protected routes here
	}

	r.Run(":8080")
}
