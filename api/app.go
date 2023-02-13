package main

import (
	"log"
	"server/api/database"
	routes "server/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	routes.RegisterRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8000"))
}

// https://github.com/gofiber/fiber/issues/270
// reference for project structure
