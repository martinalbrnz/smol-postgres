package user_routes

import (
	user_controllers "smol-api/src/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	api := app.Group("/users")

	api.Get("", user_controllers.GetUsers)
	api.Get(":id", user_controllers.GetUser)

	api.Post("", user_controllers.CreateUser)
}
