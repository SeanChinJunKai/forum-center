package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/SeanChinJunKai/forum-center/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// Function to generate a JWT token everytime a user logs in or registers
func GenerateToken(username string, email string) (token string, err error) {
	claims := &Claim{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte("secret"))
	return
}

// Function to validate JWT token
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*Claim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

// Function to get current User
func GetUserName(c *gin.Context) (string, models.ErrorResponse) {
	cookie, _ := c.Cookie("gin_cookie")

	token, err := jwt.ParseWithClaims(
		cookie,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
	)
	if err != nil {
		return "", models.ErrorResponse{Code: http.StatusUnauthorized, Message: err.Error()}
	}

	claim, ok := token.Claims.(*Claim)

	if !ok {
		return "", models.ErrorResponse{Code: http.StatusUnauthorized, Message: "couldn't parse claims"}
	}
	return claim.Username, models.ErrorResponse{}
}
