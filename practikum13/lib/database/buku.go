package database

import (
	"go-m.-azwan-alterra/practikum13/config"
	"go-m.-azwan-alterra/practikum13/model"
)

func GetBuku() (interface{}, error) {
	var bukus []model.Buku

	if e := config.DB.Find(&bukus).Error; e != nil {
		return nil, e
	}

	return bukus, nil
}
