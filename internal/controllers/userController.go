package controllers

import (
	"fmt"
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/models"
	"github.com/SeanChinJunKai/forum-center/internal/services"
	"github.com/SeanChinJunKai/forum-center/internal/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	fmt.Println("Registering User")
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err2 := services.RegisterUser(input)

	if err2.Message != "" {
		c.AbortWithStatusJSON(err2.Code, gin.H{"error": err2.Message})
		return
	}

	token, err := utils.GenerateToken(user.Name, user.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("gin_cookie", token, 86400, "/", "localhost", true, true)
	c.JSON(http.StatusCreated, gin.H{"name": user.Name, "email": user.Email})
}

// Function that allows user to login based on user input of username and password
func LoginUser(c *gin.Context) {
	fmt.Println("User Login")
	var input models.LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err2 := services.LoginUser(input)

	if err2.Message != "" {
		c.AbortWithStatusJSON(err2.Code, gin.H{"error": err2.Message})
		return
	}

	token, err := utils.GenerateToken(user.Name, user.Email)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("gin_cookie", token, 86400, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{"name": user.Name, "email": user.Email})

}
