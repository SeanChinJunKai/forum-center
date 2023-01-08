package controllers

import (
	"fmt"
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/models"
	"github.com/SeanChinJunKai/forum-center/internal/services"
	"github.com/SeanChinJunKai/forum-center/internal/utils"
	"github.com/gin-gonic/gin"
)

// Function to allow creation of comments
func CreateComment(c *gin.Context) {
	fmt.Println("Creating Comment")
	var input models.CreateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currUser, err := utils.GetUserName(c)
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	comment, err := services.CreateComment(input, currUser)
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": comment})

}

// Function to allow updating of comments
func UpdateComment(c *gin.Context) {
	fmt.Println("Updating Specified Comment")
	var input models.UpdateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currUser, err := utils.GetUserName(c)
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	comment, err := services.UpdateComment(input, currUser, c.Param("commentId"))
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comment})

}

// Function to allow deleting of comments
func DeleteComment(c *gin.Context) {
	fmt.Println("Deleting Specified Comment")
	currUser, err := utils.GetUserName(c)
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	response, err := services.DeleteComment(currUser, c.Param("commentId"))
	if err.Message != "" {
		c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}
