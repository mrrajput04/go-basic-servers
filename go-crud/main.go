package main

import (
	"crud/controllers"
	"crud/model"

	"github.com/gin-gonic/gin" // You need to run: go get -u github.com/gin-gonic/gin
)

func main() {
	router := gin.Default()

	model.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.PATCH("posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run("localhost:8000")
}
