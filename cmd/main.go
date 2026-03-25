package main

import (
	"go-backend/config"
	"go-backend/migration"
	"go-backend/routes"
)

func main() {
	config.InitDB()
	migration.Migrate()

	r := routes.SetupRoutes()
	r.Run(":8080")
}