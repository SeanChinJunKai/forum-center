package services

import (
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/dataaccess"
	"github.com/SeanChinJunKai/forum-center/internal/models"
)

func RegisterUser(input models.CreateUserInput) (models.User, models.ErrorResponse) {
	if input.Name == "" {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Name field is empty"}
	}

	if input.Email == "" {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Email field is empty"}
	}

	if input.Password == "" {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Password field is empty"}
	}

	return dataaccess.RegisterUser(input.Name, input.Email, input.Password)

}

func LoginUser(input models.LoginUserInput) (models.User, models.ErrorResponse) {
	if input.Name == "" {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Name field is empty"}
	}

	if input.Password == "" {
		return models.User{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Password field is empty"}
	}

	return dataaccess.LoginUser(input.Name, input.Password)
}
