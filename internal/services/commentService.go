package services

import (
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/dataaccess"
	"github.com/SeanChinJunKai/forum-center/internal/models"
)

func CreateComment(input models.CreateCommentInput, username string) (models.Comment, models.ErrorResponse) {
	if input.Content == "" {
		return models.Comment{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Please enter description for your comment"}
	}
	return dataaccess.CreateComment(input.Content, input.PostID, input.Reference, username)
}

func UpdateComment(input models.UpdateCommentInput, username string, commentId string) (models.Comment, models.ErrorResponse) {
	if input.Like {
		return dataaccess.UpdateCommentLikes(username, commentId)
	}

	if input.Dislike {
		return dataaccess.UpdateCommentDislikes(username, commentId)
	}

	if input.Content != "" {
		return dataaccess.UpdateCommentContent(input.Content, username, commentId)
	}

	return models.Comment{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Please add description"}
}

func DeleteComment(username string, commentId string) (string, models.ErrorResponse) {
	return dataaccess.DeleteComment(username, commentId)
}
