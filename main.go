package main

import (
	"log"
	"os"
	"smol-api/src/db"
	"smol-api/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	uri := os.Getenv("DB_URI")

	if port == "" {
		port = "8080"
	}

	db := db.GetClient(uri)

	// Setup app
	app := fiber.New(fiber.Config{EnablePrintRoutes: true})

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.RegisterRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":" + port))
}
