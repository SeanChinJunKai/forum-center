package models

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Tags      string        `json:"tags"`
	Likes     []PostLike    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"postlikes"`
	Dislikes  []PostDislike `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"postdislikes"`
	Comments  []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	UserID    uint          // foreign key
}

type CreatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	Tags    string `json:"tags"`
}

type UpdatePostInput struct {
	Like    bool   `json:"like"`
	Dislike bool   `json:"dislike"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}
