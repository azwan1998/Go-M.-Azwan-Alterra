package appModel

import (
	"gorm.io/gorm"
)

type UserDbModel struct {
	db *gorm.DB
}

func NewUserDbModel(db *gorm.DB) *UserDbModel {
	// db.AutoMigrate(&User{})
	return &UserDbModel{
		db: db,
	}
}

func (pm *UserDbModel) GetByEmailAndPassword(email string, password string) (User, error) {
	p := User{}
	err := pm.db.Where("email = ? AND password = ?", email, password).First(&p).Error
	return p, err
}

func (pm *UserDbModel) GetAll() ([]User, error) {
	var allUser []User
	err := pm.db.Find(&allUser).Error
	return allUser, err
}

func (pm *UserDbModel) Add(p User) (User, error) {
	err := pm.db.Save(&p).Error
	return p, err
}
