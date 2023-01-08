package models

import (
	"time"
)

type Comment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string           `json:"content"`
	PostID    uint             `json:"postid"`    // foreign key
	Reference uint             `json:"reference"` // the commment u replying to
	Likes     []CommentLike    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"commentlikes"`
	Dislikes  []CommentDislike `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"commentdislikes"`
	UserID    uint
}

type CreateCommentInput struct {
	Content   string `json:"content"`
	PostID    uint   `json:"postid"`
	Reference uint   `json:"reference"`
}

type UpdateCommentInput struct {
	Like    bool   `json:"like"`
	Dislike bool   `json:"dislike"`
	Content string `json:"content"`
}
