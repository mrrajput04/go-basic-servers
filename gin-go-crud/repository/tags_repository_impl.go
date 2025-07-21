package repository

import (
	"errors"
	"gin-crud/data/request"
	"gin-crud/helper"
	"gin-crud/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTags = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) FindById(tagsId int) (model.Tags, error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	results := t.Db.Find(&tags)
	helper.ErrorPanic(results.Error)
	return tags
}
