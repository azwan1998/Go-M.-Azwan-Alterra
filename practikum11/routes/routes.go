package routes

import (
	"go-m.-azwan-alterra/practikum11/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Routing
	e.GET("/books", controller.GetBooks)
	e.POST("/create-books", controller.CreateBooks)

	return e
}
