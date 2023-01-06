package models

import (
	"time"
)

type Comment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Content   string `json:"content"`
	PostID    uint   `json:"postid"` // foreign key
	MainID    *uint
	Replies   []Comment        `gorm:"foreignkey:MainID" constraint:"OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Likes     []CommentLike    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"commentlikes"`
	Dislike   []CommentDislike `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"commentdislikes"`
	UserID    int
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CreateCommentInput struct {
	Content   string `json:"content"`
	PostID    uint   `json:"postid"`
	Reference int    `json:"reference"`
}

type UpdateCommentInput struct {
	Content string `json:"content"`
}
