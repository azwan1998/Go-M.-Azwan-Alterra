package controller

import (
	"net/http"

	"go-learn-alteraa/config"
	"go-learn-alteraa/model"

	"github.com/labstack/echo/v4"
)

func GetBooks(context echo.Context) error {
	var buku []model.Buku

	if err := config.DB.Find(&buku).Error; err != nil {
		return err
	}
	return context.JSON(http.StatusOK, buku)
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
