package routes

import (
	"blog-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("", handlers.CreatePost)
		postRoutes.GET("", handlers.GetPosts)
		postRoutes.GET("/:id", handlers.GetPostByID)
		postRoutes.PUT("/:id", handlers.UpdatePost)
		postRoutes.DELETE("/:id", handlers.DeletePost)
	}
}
