package routes

import (
	"go-m.-azwan-alterra/practikum12/constants"
	"go-m.-azwan-alterra/practikum12/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Routing
	e.GET("/books", controller.GetBooks)
	e.POST("/create-books", controller.CreateBooks)
	// e.POST("/login", controller.LoginUsersController)

	//JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/create", controller.CreateUserController)
	r.GET("/users", controller.GetUserController)

	return e
}
