package auth_routes

import (
	auth_controllers "smol-api/src/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	api := app.Group("/auth")

	api.Post("", auth_controllers.AuthUser)
}
