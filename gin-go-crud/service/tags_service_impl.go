package service

import (
	"gin-crud/data/request"
	"gin-crud/data/response"
	"gin-crud/helper"
	"gin-crud/model"
	"gin-crud/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

func (t TagsServiceImpl) Create(tag request.CreateTagsRequest) {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tag.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t TagsServiceImpl) Update(tag request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tag.Id)
	helper.ErrorPanic(err)
	tagData.Name = tag.Name
	t.TagsRepository.Update(tagData)
}

func (t TagsServiceImpl) Delete(tagId int) {
	t.TagsRepository.Delete(tagId)
}

func (t TagsServiceImpl) FindById(tagId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse
}

func (t TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}
