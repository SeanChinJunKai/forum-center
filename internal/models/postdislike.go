package models

type PostDislike struct {
	ID     uint `gorm:"primarykey"`
	PostID uint // foreign key
	UserID int
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
