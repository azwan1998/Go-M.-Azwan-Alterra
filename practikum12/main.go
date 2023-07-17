package main

import (
	"go-m.-azwan-alterra/practikum12/config"
	"go-m.-azwan-alterra/practikum12/middlewares"
	"go-m.-azwan-alterra/practikum12/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
