package database

import (
	"go-m.-azwan-alterra/practikum13/config"
	"go-m.-azwan-alterra/practikum13/model"
)

func GetUsers() (interface{}, error) {
	var users []model.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func CreateUser(user *model.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
func GetDetailUsers(userId int) (interface{}, error) {
	var user model.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

// func LoginUsers(user *model.User) (interface{}, error) {

// 	var err error
// 	if err = config.DB.Where("email = ? AND password = ?",
// 		user.Email, user.Password).First(user).Error; err != nil {
// 		return nil, err
// 	}

// 	user.Token, err = middlewares.CreateToken(int(user.ID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := config.DB.Save(user).Error; err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
