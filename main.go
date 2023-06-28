package main

import (
	"go-learn-alteraa/config"
	"go-learn-alteraa/router"
)

func main() {
	config.InitDB()
	// InitialMigration()
	e := router.New()
	e.Start(":8000")
}
