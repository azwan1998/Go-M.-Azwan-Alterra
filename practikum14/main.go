package main

import (
	"belajar-go-echo/appConfig"
	"belajar-go-echo/appController"
	"belajar-go-echo/appMiddleware"
	"belajar-go-echo/appModel"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := appConfig.NewConfig()
	if err != nil {
		panic(err)
	}

	// UserModel can be either UserMemModel or UserDbModel, depends on the configuration
	var UserModel appModel.UserModel
	switch cfg.Storage {
	case "db":
		db, err := gorm.Open(mysql.Open(cfg.ConnectionString), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		UserModel = appModel.NewUserDbModel(db)

		// create new echo instant
		e := echo.New()
		appMiddleware.AddGlobalMiddlewares(e)
		appController.HandleRoutes(e, cfg.JwtSecret, UserModel)

		if err = e.Start(fmt.Sprintf(":%d", cfg.HttpPort)); err != nil {
			panic(err)
		
	}
}
