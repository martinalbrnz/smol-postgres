package routes_accounts

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hola accounts")
}

func Register(app *fiber.App) {
	app.Get("/accounts", Hello)
}
