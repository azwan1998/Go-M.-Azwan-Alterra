package appModel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}
