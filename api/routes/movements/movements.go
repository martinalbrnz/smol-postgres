package routes_movements

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hola movements")
}

func HelloId(c *fiber.Ctx) error {
	return c.SendString("Hola movimiento " + c.Params("id"))
}

func Register(app *fiber.App) {
	movements := app.Group("/movements")
	movements.Get("/", Hello)
	movements.Get("/:id", HelloId)
}
