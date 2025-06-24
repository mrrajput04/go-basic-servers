package main

import (
	"crud/controllers"
	"crud/models"

	"github.com/gin-gonic/gin" // You need to run: go get -u github.com/gin-gonic/gin
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.FindPost)

	router.Run("localhost:8000")
}
