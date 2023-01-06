package models

type PostLike struct {
	ID     uint `gorm:"primarykey"`
	PostID uint // foreign key
	UserID int
	User   User
}
