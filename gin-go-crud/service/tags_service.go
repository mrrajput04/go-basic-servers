package service

import (
	"gin-crud/data/request"
	"gin-crud/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsid int) response.TagsResponse
	FindAll() []response.TagsResponse
}
