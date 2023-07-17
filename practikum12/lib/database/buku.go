package database

import (
	"go-m.-azwan-alterra/practikum12/config"
	"go-m.-azwan-alterra/practikum12/model"
)

func GetBuku() (interface{}, error) {
	var bukus []model.Buku

	if e := config.DB.Find(&bukus).Error; e != nil {
		return nil, e
	}

	return bukus, nil
}
