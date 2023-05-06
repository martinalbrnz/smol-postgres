package routes

import (
	auth_routes "smol-api/src/routes/auth"
	user_routes "smol-api/src/routes/users"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	user_routes.Register(app)
	auth_routes.Register(app)
}
