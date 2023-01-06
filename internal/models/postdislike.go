package models

type PostDislike struct {
	ID     uint `gorm:"primarykey"`
	PostID uint // foreign key
	UserID uint
}
