package models

type CommentDislike struct {
	ID        uint `gorm:"primarykey"`
	CommentID uint
	UserID    int
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
