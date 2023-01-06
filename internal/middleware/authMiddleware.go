package middleware

import (
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/utils"
	"github.com/gin-gonic/gin"
)

// Authentication middleware protects routes which requires a user to be logged in
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("gin_cookie")
		c.Request.Header.Set("Authorization", cookie)
		splitToken := c.Request.Header.Get("Authorization")
		if splitToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, no token!"})
			return
		}

		err := utils.ValidateToken(splitToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}
