package dataaccess

import (
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/config"
	"github.com/SeanChinJunKai/forum-center/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, useremail string, userpassword string) (models.User, models.ErrorResponse) {
	var checkUser models.User
	config.DB.Where("name=?", username).Find(&checkUser)
	if checkUser.ID != 0 {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Username already exists. Please choose another username"}
	}
	config.DB.Where("email=?", useremail).Find(&checkUser)
	if checkUser.ID != 0 {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Email already registered"}
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userpassword), 14)

	user := models.User{
		Name:     username,
		Email:    useremail,
		Password: password,
	}
	config.DB.Create(&user)

	return user, models.ErrorResponse{}
}

func LoginUser(username string, userpassword string) (models.User, models.ErrorResponse) {
	var user models.User
	config.DB.Where("name=?", username).Find(&user)
	if user.ID == 0 {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "No such user exists"}
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(userpassword)); err != nil {
		return models.User{}, models.ErrorResponse{Code: http.StatusUnauthorized, Message: "Incorrect password"}
	}

	return user, models.ErrorResponse{}
}
