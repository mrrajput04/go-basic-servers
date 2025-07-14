package repository

import (
	"errors"
	"gin-crud/data/request"
	"gin-crud/helper"
	"gin-crud/models"

	"gorm.io/gorm"
)

type TagsRepository interface {
	Save(tags models.Tags)
	Update(tags models.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags models.Tags, err error)
	FindAll() []models.Tags
}

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t TagsRepositoryImpl) Save(tags models.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) Update(tags models.Tags) {
	var updateTags = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) Delete(tagsId int) {
	var tags models.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) FindById(tagsId int) (models.Tags, error) {
	var tag models.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t TagsRepositoryImpl) FindAll() []models.Tags {
	var tags []models.Tags
	results := t.Db.Find(&tags)
	helper.ErrorPanic(results.Error)
	return tags
}
