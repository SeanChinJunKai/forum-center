package models

type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
	Posts    []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"posts"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
