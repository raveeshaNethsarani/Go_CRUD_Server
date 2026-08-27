package handlers

import (
	"net/http"
	"blog-api/config"
	"blog-api/models"
	"github.com/gin-gonic/gin"
)

// ----------------------CreatePost handles POST /posts.----------------------
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

//---------------------- GetPosts handles GET /posts.----------------------
func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := config.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

//----------------------GetPostByID	handles	GET	/posts/:id----------------------
func GetPostByID(c	*gin.Context) {
id	:=	c.Param("id")
var	post	models.Post
result	:=	config.DB.First(&post,	id)
if	result.Error	!=	nil	{
c.JSON(http.StatusNotFound,	gin.H{"error":	"Post	not	found"})
return
}
c.JSON(http.StatusOK, post)}

///---------------------- UpdatePost handles PUT /posts/:id----------------------
func UpdatePost(c *gin.Context) {
id := c.Param("id")
var post models.Post
// First, find the existing post
if err := config.DB.First(&post, id).Error; err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
return
}
// Bind the new data from request body
var input models.Post
if err := c.ShouldBindJSON(&input); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}
// Update fields
post.Title = input.Title
post.Content = input.Content
if err := config.DB.Save(&post).Error; err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}
c.JSON(http.StatusOK, post)
}

///---------------------- DeletePost handles DELETE /posts/:id----------------------
func DeletePost(c *gin.Context) {
id := c.Param("id")
var post models.Post
// Check if the post exists first
if err := config.DB.First(&post, id).Error; err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
return
}
if err := config.DB.Delete(&post).Error; err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}
c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
