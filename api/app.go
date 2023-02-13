package main

import (
	"log"
	"server/api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	// app.Get("/movements", routes)
}

func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola mundo!")
	})

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	// movementRoute := app.Group("/movements")

	log.Fatal(app.Listen(":8000"))
}
