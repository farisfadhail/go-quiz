package main

import (
	"github.com/gofiber/fiber/v2"
	"go-quiz/database"
	"go-quiz/database/migrations"
	"go-quiz/routes"
	"log"
)

func main() {
	// Database Init
	database.DatabaseInit()
	migrations.RunMigration()

	app := fiber.New()

	routes.RouteInit(app)

	err := app.Listen("localhost:3000")
	if err != nil {
		log.Println("Failed to listen Go Fiber Server")
	}
}
