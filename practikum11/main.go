package main

import (
	"go-m.-azwan-alterra/practikum11/config"
	"go-m.-azwan-alterra/practikum11/routes"
)

func main() {
	config.InitDB()
	// InitialMigration()
	e := routes.New()
	e.Start(":8000")
}
