package main

import (
	"go-m.-azwan-alterra/practikum13/config"
	"go-m.-azwan-alterra/practikum13/middlewares"
	"go-m.-azwan-alterra/practikum13/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
