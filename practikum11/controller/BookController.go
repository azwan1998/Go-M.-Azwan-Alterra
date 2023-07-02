package controller

import (
	"net/http"

	"go-m.-azwan-alterra/practikum11/config"
	"go-m.-azwan-alterra/practikum11/lib/database"
	"go-m.-azwan-alterra/practikum11/model"

	"github.com/labstack/echo/v4"
)

func GetBooks(context echo.Context) error {
	bukus, e := database.GetBuku()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return context.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  bukus,
	})
}

func CreateBooks(context echo.Context) error {
	var params model.Buku
	err := context.Bind(&params)
	if err != nil {
		return err
	}

	if err := config.DB.Create(&params).Error; err != nil {
		return err
	}

	return context.JSON(http.StatusCreated, params)
}

func Show(id string) (*model.Buku, error) {
	var buku model.Buku
	result := config.DB.First(&buku, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &buku, nil
}
