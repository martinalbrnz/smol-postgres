package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// app.Get("/movements", routes.)
}

func main() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola mundo!")
	})

	// movementRoute := app.Group("/movements")

	app.Listen(":8000")
}
