package main

import (
	"github.com/AntoniusIvan/div45rhino-trivia-main/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// app := fiber.New()

	// app.Listen(":3000")

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, GrunshvdfgdfSD:")
	})

	app.Listen(":3000")
}
