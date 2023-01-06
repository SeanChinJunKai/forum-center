package config

import (
	"fmt"

	"github.com/SeanChinJunKai/forum-center/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable to allow other files to access database

// Function to connect to MySQL Database
func Connect() {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/forumcenter?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Successful")
	DB = connection
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Comment{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.PostLike{})
	DB.AutoMigrate(&models.PostDislike{})
	DB.AutoMigrate(&models.CommentDislike{})
	DB.AutoMigrate(&models.CommentLike{})
}
