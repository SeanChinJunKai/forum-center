package models

type CommentDislike struct {
	ID        uint `gorm:"primarykey"`
	CommentID uint
	UserID    uint
}
