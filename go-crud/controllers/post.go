package controllers

import (
	"crud/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateExistingPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := model.Post{Title: input.Title, Content: input.Content}
	model.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})

}

func FindPosts(c *gin.Context) {
	var posts []model.Post
	model.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"Data": posts})

}

func FindPost(c *gin.Context) {
	var post model.Post

	if err := model.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": post})
}

func UpdatePost(c *gin.Context) {
	var post model.Post
	if err := model.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	var input UpdateExistingPost

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatePost := model.Post{Title: input.Title, Content: input.Content}

	model.DB.Model(&post).Updates(updatePost)
	c.JSON(http.StatusOK, gin.H{"data": post})

}

func DeletePost(c *gin.Context) {
	var post model.Post
	if err := model.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}
	model.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"Data": "Post deleted successfully"})
}
