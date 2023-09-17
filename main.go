package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"test-app/database"
	"test-app/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	go database.Init()
	database.Connect()

	routes.Setup(app)

	app.Listen(":8080")
}
