package appController

import (
	"belajar-go-echo/appModel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRoutes(e *echo.Echo, jwtSecret string, UserModel appModel.UserModel, profileModel appModel.ProfileModel) UserController {

	UserController := NewUserController(jwtSecret, UserModel, profileModel)

	e.POST("/login", UserController.Login)
	e.POST("/login/", UserController.Login)

	e.POST("/add", UserController.Add)
	e.POST("/add/", UserController.Add)

	jwtMiddleware := middleware.JWT([]byte(jwtSecret))

	//users
	e.GET("/users", UserController.GetAll, jwtMiddleware)
	e.GET("/users/", UserController.GetAll, jwtMiddleware)

	return UserController
}
