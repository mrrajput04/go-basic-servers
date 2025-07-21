package controller

import (
	"gin-crud/data/request"
	"gin-crud/data/response"
	"gin-crud/helper"
	"gin-crud/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TagController struct {
	tagService service.TagsService
}

func NewTagsController(service service.TagsService) *TagController {
	return &TagController{
		tagService: service,
	}
}

func (controller *TagController) Create(ctx *gin.Context) {
	log.Info().Msg("create tags")
	createTagRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.tagService.Create(createTagRequest)

	webResponse := response.Response{
		Code:   201,
		Status: "Created",
		Data:   nil,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *TagController) Update(ctx *gin.Context) {
	log.Info().Msg("update tags")
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagRequest.Id = id

	controller.tagService.Update(updateTagRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete tags")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagService.Delete(id)

	webResponse := response.Response{
		Status: "Ok",
		Code:   200,
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindById(ctx *gin.Context) {
	log.Info().Msg("get by id tags")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}

	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindAll(ctx *gin.Context) {
	log.Info().Msg("find all tags")
	tagResponse := controller.tagService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}

	ctx.Header("content-type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
