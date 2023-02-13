package routes

import (
	routes_accounts "server/api/routes/accounts"
	routes_movements "server/api/routes/movements"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello from routes!")
}

func RegisterRoutes(app *fiber.App) {
	routes_accounts.Register(app)
	routes_movements.Register(app)
}
