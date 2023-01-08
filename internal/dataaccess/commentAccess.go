package dataaccess

import (
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/config"
	"github.com/SeanChinJunKai/forum-center/internal/models"
	"gorm.io/gorm/clause"
)

func CreateComment(content string, postId uint, reference uint, username string) (models.Comment, models.ErrorResponse) {
	var author models.User
	config.DB.Where("name=?", username).Find(&author)

	var comment models.Comment

	if reference == 0 {
		comment = models.Comment{
			// no main reference
			Content: content,
			PostID:  postId,
			UserID:  author.ID,
		}
		config.DB.Create(&comment)
		config.DB.Preload(clause.Associations).Find(&comment, &comment)
		return comment, models.ErrorResponse{}
	}
	comment = models.Comment{
		Content:   content,
		PostID:    postId,
		Reference: reference,
		UserID:    author.ID,
	}
	config.DB.Create(&comment)
	config.DB.Preload(clause.Associations).Find(&comment, &comment)
	return comment, models.ErrorResponse{}

}

func UpdateCommentLikes(username string, commentId string) (models.Comment, models.ErrorResponse) {
	var specifiedComment models.Comment
	config.DB.Where("id=?", commentId).Find(&specifiedComment)
	if specifiedComment.ID == 0 {
		return models.Comment{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified comment does not exist"}
	}
	var author models.User
	config.DB.Where("name=?", username).Find(&author)

	var updatedComment models.Comment
	var findLike models.CommentLike
	config.DB.Where("comment_id=?", commentId).Where("user_id=?", author.ID).Find(&findLike)
	if findLike.ID != 0 {
		config.DB.Unscoped().Delete(&findLike)
		config.DB.Preload(clause.Associations).Find(&updatedComment, &specifiedComment)

	} else {
		like := models.CommentLike{
			CommentID: specifiedComment.ID,
			UserID:    uint(author.ID),
		}
		config.DB.Create(&like)
		config.DB.Preload(clause.Associations).Find(&updatedComment, &specifiedComment)
	}
	return updatedComment, models.ErrorResponse{}
}

func UpdateCommentDislikes(username string, commentId string) (models.Comment, models.ErrorResponse) {
	var specifiedComment models.Comment
	config.DB.Where("id=?", commentId).Find(&specifiedComment)
	if specifiedComment.ID == 0 {
		return models.Comment{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified comment does not exist"}
	}
	var author models.User
	config.DB.Where("name=?", username).Find(&author)
	var updatedComment models.Comment
	var findDislike models.CommentDislike
	config.DB.Where("comment_id=?", commentId).Where("user_id=?", author.ID).Find(&findDislike)

	if findDislike.ID != 0 {
		config.DB.Unscoped().Delete(&findDislike)
		config.DB.Preload(clause.Associations).Find(&updatedComment, &specifiedComment)

	} else {
		dislike := models.CommentDislike{
			UserID:    uint(author.ID),
			CommentID: specifiedComment.ID,
		}
		config.DB.Create(&dislike)
		config.DB.Preload(clause.Associations).Find(&updatedComment, &specifiedComment)
	}
	return updatedComment, models.ErrorResponse{}
}

func UpdateCommentContent(content string, username string, commentId string) (models.Comment, models.ErrorResponse) {
	var specifiedComment models.Comment
	config.DB.Where("id=?", commentId).Find(&specifiedComment)
	if specifiedComment.ID == 0 {
		return models.Comment{}, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified comment does not exist"}
	}
	var author models.User
	config.DB.Where("name=?", username).Find(&author)
	var updatedComment models.Comment
	if author.ID != specifiedComment.UserID {
		return models.Comment{}, models.ErrorResponse{Code: http.StatusUnauthorized, Message: "Current user is not the author of this comment"}
	}
	specifiedComment.Content = content
	config.DB.Save(&specifiedComment)
	config.DB.Preload(clause.Associations).Find(&updatedComment, &specifiedComment)
	return updatedComment, models.ErrorResponse{}

}

func DeleteComment(username string, commentId string) (string, models.ErrorResponse) {
	var comment models.Comment
	config.DB.Where("id=?", commentId).Find(&comment)

	if comment.ID == 0 {
		return "", models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified comment does not exist"}
	}

	var author models.User
	config.DB.Where("name=?", username).Find(&author)
	if comment.UserID != author.ID {
		return "", models.ErrorResponse{Code: http.StatusUnauthorized, Message: "Current user is not the author of this post"}
	}

	config.DB.Model(models.Comment{}).Where("reference=?", commentId).Update("reference", 0)
	config.DB.Unscoped().Delete(&comment)
	return commentId, models.ErrorResponse{}
}
