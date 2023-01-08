package models

type CommentLike struct {
	ID        uint `gorm:"primarykey"`
	CommentID uint
	UserID    uint
}
