package services

import (
	"fmt"
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/dataaccess"
	"github.com/SeanChinJunKai/forum-center/internal/models"
)

func GetPosts() []models.Post {
	return dataaccess.GetPosts()
}

func GetPostById(postId string) (models.Post, models.ErrorResponse) {
	return dataaccess.GetPostById(postId)
}

func CreatePost(input models.CreatePostInput, username string) (models.Post, models.ErrorResponse) {
	if input.Title == "" {
		return models.Post{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Please add a title"}
	}

	if input.Content == "" {
		return models.Post{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Please add a description"}
	}

	if input.Tags == "" {
		return models.Post{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Please add tags"}
	}

	return dataaccess.CreatePost(input.Title, input.Content, input.Tags, username)
}

func UpdatePost(input models.UpdatePostInput, username string, postId string) (models.Post, models.ErrorResponse) {
	fmt.Println(input.Like)
	if input.Like {
		return dataaccess.UpdatePostLikes(username, postId)
	}

	if input.Dislike {
		return dataaccess.UpdatePostDislikes(username, postId)
	}

	if input.Content != "" {
		return dataaccess.UpdatePostContent(input.Content, username, postId)
	}

	if input.Tags != "" {
		return dataaccess.UpdatePostTags(input.Tags, username, postId)
	}

	if input.Content == "" {
		return models.Post{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Please add description"}
	}
	return models.Post{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Please add tags"}
}

func DeletePost(username string, postId string) (string, models.ErrorResponse) {
	return dataaccess.DeletePost(username, postId)
}
