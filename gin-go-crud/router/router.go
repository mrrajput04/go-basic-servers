package router

import (
	"gin-crud/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func newRouter(tagController *controller.TagController) *gin.Engine {

	r := gin.Default()

	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello world")
	})

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	router := r.Group("/api")
	tagRouter := router.Group("/tag")
	tagRouter.GET("", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PATCH("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	return r
}
