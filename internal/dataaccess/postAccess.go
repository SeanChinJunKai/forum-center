package dataaccess

import (
	"net/http"

	"github.com/SeanChinJunKai/forum-center/internal/config"
	"github.com/SeanChinJunKai/forum-center/internal/models"
	"gorm.io/gorm/clause"
)

func GetPosts() []models.Post {
	var posts []models.Post
	config.DB.Model(&models.Post{}).Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&posts)
	return posts
}

func GetPostById(postId string) (models.Post, models.ErrorResponse) {
	var post models.Post
	config.DB.Where("id=?", postId).Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&post)
	if post.ID == 0 {
		return post, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified post does not exist"}
	}
	return post, models.ErrorResponse{}
}

func CreatePost(title string, content string, tags string, username string) (models.Post, models.ErrorResponse) {

	var author models.User
	config.DB.Where("name=?", username).Find(&author)

	post := models.Post{
		Title:   title,
		Content: content,
		Tags:    tags,
		UserID:  author.ID,
	}
	config.DB.Create(&post)
	config.DB.Preload(clause.Associations).Find(&post, &post)
	return post, models.ErrorResponse{}
}

func UpdatePostLikes(username string, postId string) (models.Post, models.ErrorResponse) {
	var specifiedPost models.Post
	config.DB.Where("id=?", postId).Find(&specifiedPost)
	if specifiedPost.ID == 0 {
		return specifiedPost, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified post does not exist"}
	}
	var author models.User
	config.DB.Where("name=?", username).Find(&author)

	var updatedPost models.Post
	var findLike models.PostLike
	config.DB.Where("post_id=?", postId).Where("user_id=?", author.ID).Find(&findLike)
	if findLike.ID != 0 {
		config.DB.Unscoped().Delete(&findLike)
		config.DB.Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&updatedPost, &specifiedPost)

	} else {
		like := models.PostLike{
			UserID: uint(author.ID),
			PostID: specifiedPost.ID,
		}
		config.DB.Create(&like)
		config.DB.Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&updatedPost, &specifiedPost)
	}
	return updatedPost, models.ErrorResponse{}
}

func UpdatePostDislikes(username string, postId string) (models.Post, models.ErrorResponse) {
	var specifiedPost models.Post
	config.DB.Where("id=?", postId).Find(&specifiedPost)
	if specifiedPost.ID == 0 {
		return specifiedPost, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified post does not exist"}
	}
	var author models.User
	config.DB.Where("name=?", username).Find(&author)
	var updatedPost models.Post
	var findDislike models.PostDislike
	config.DB.Where("post_id=?", postId).Where("user_id=?", author.ID).Find(&findDislike)
	if findDislike.ID != 0 {
		config.DB.Unscoped().Delete(&findDislike)
		config.DB.Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&updatedPost, &specifiedPost)

	} else {
		dislike := models.PostDislike{
			UserID: uint(author.ID),
			PostID: specifiedPost.ID,
		}
		config.DB.Create(&dislike)
		config.DB.Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&updatedPost, &specifiedPost)
	}
	return updatedPost, models.ErrorResponse{}
}

func UpdatePostContent(content string, username string, postId string) (models.Post, models.ErrorResponse) {
	var specifiedPost models.Post
	config.DB.Where("id=?", postId).Find(&specifiedPost)
	if specifiedPost.ID == 0 {
		return specifiedPost, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified post does not exist"}
	}
	var author models.User
	config.DB.Where("name=?", username).Find(&author)
	var updatedPost models.Post
	if author.ID != specifiedPost.UserID {
		return models.Post{}, models.ErrorResponse{Code: http.StatusUnauthorized, Message: "Current user is not the author of this post"}
	}
	specifiedPost.Content = content
	config.DB.Save(&specifiedPost)
	config.DB.Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&updatedPost, &specifiedPost)
	return updatedPost, models.ErrorResponse{}
}

func UpdatePostTags(tags string, username string, postId string) (models.Post, models.ErrorResponse) {
	var specifiedPost models.Post
	config.DB.Where("id=?", postId).Find(&specifiedPost)
	if specifiedPost.ID == 0 {
		return specifiedPost, models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified post does not exist"}
	}
	var author models.User
	config.DB.Where("name=?", username).Find(&author)
	var updatedPost models.Post
	if author.ID != specifiedPost.UserID {
		return models.Post{}, models.ErrorResponse{Code: http.StatusUnauthorized, Message: "Current user is not the author of this post"}
	}
	specifiedPost.Tags = tags
	config.DB.Save(&specifiedPost)
	config.DB.Preload("Comments.Likes").Preload("Comments.Dislikes").Preload(clause.Associations).Find(&updatedPost, &specifiedPost)
	return updatedPost, models.ErrorResponse{}
}

func DeletePost(username string, postId string) (string, models.ErrorResponse) {
	var post models.Post
	config.DB.Where("id=?", postId).Find(&post)

	if post.ID == 0 {
		return "", models.ErrorResponse{Code: http.StatusBadRequest, Message: "Specified post does not exist"}
	}

	var author models.User
	config.DB.Where("name=?", username).Find(&author)
	if post.UserID != author.ID {
		return "", models.ErrorResponse{Code: http.StatusUnauthorized, Message: "Current user is not the author of this post"}
	}

	config.DB.Unscoped().Delete(&post)
	return postId, models.ErrorResponse{}
}
