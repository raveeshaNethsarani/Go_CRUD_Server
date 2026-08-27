package main

import (
	"blog-api/config" //import config package
	"blog-api/models" //import models package
	"blog-api/routes" //import routes package
	"github.com/gin-gonic/gin" //import gin package
)

func main() {
	config.ConnectDatabase() //connect to the database
	//migrate the Post model to the database
	config.DB.AutoMigrate(&models.Post{})

	r := gin.Default()
	routes.SetupRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
