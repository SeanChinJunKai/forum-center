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
	Replies   []Comment        `gorm:"foreignkey:MainID" constraint:"OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Likes     []CommentLike    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commentlikes"`
	Dislike   []CommentDislike `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"commentdislikes"`
	UserID    int
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateCommentInput struct {
	Content   string `json:"content"`
	PostID    uint   `json:"postid"`
	Reference int    `json:"reference"`
}

type UpdateCommentInput struct {
	Content string `json:"content"`
}
