package controllers

import (
	"fmt"
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/models"
	"github.com/SeanChinJunKai/forum-center/internal/services"
	"github.com/SeanChinJunKai/forum-center/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	fmt.Println("Getting All Posts")
	posts := services.GetPosts()
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetPostById(c *gin.Context) {
	fmt.Println("Getting Specified Post")
	post, err := services.GetPostById(c.Param("postId"))
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func CreatePost(c *gin.Context) {
	fmt.Println("Creating Post")
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currUser, err := utils.GetUserName(c)
	if err.Message != "" {
		fmt.Println(err.Message)
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	post, err := services.CreatePost(input, currUser)
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": post})
}

func UpdatePost(c *gin.Context) {
	fmt.Println("Updating Post")
	var input models.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currUser, err := utils.GetUserName(c)
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	post, err := services.UpdatePost(input, currUser, c.Param("postId"))
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

// Function that allows a user to delete a post given its ID only if the user is the author of the post
func DeletePost(c *gin.Context) {
	fmt.Println("Deleting Specified Post")
	currUser, err := utils.GetUserName(c)
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	response, err := services.DeletePost(currUser, c.Param("postId"))
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}
