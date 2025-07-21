package main

import (
	"gin-crud/config"
	"gin-crud/helper"
	"gin-crud/model"
	"gin-crud/repository"
	"gin-crud/router"
	"gin-crud/service"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
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

	//router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:           ":8888",
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
