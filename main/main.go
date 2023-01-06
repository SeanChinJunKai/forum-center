package main

import (
	"github.com/SeanChinJunKai/forum-center/internal/config"
	"github.com/SeanChinJunKai/forum-center/internal/controllers"
	"github.com/SeanChinJunKai/forum-center/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Connect() // Connect to the MySQL Database

	// Route setup
	router := gin.Default()
	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", controllers.RegisterUser)
			users.POST("/login", controllers.LoginUser)
			users.POST("/logout", controllers.LogoutUser)
		}
		post := api.Group("/post")
		{
			post.GET("/", controllers.GetPosts)
			post.POST("/", middleware.Auth(), controllers.CreatePost)
			post.GET("/:postId", controllers.GetPostById)
			post.PUT("/:postId", middleware.Auth(), controllers.UpdatePost)
			post.DELETE("/:postId", middleware.Auth(), controllers.DeletePost)

		}
	}

	router.Run(":8080")

}
