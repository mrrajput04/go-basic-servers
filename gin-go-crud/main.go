package main

import (
	"gin-crud/config"
	"gin-crud/controller"
	"gin-crud/helper"
	"gin-crud/model"
	"gin-crud/repository"
	"gin-crud/router"
	"gin-crud/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	log.Info().Msg("Started Server!")

	//database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//init repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	//init service
	tagsService := service.NewTagServiceImpl(tagsRepository, validate)

	//init controller
	tagsController := controller.NewTagsController(tagsService)

	r := gin.Default()
	url := ginSwagger.URL("http://localhost:8080/swagger/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		helper.ErrorPanic(err)
	}
}
