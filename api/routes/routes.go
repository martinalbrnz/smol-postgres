package routes

import (
	"server/api/database"
	"server/api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type dataResponse struct {
	id          string
	date        time.Time
	account     string
	amount      float32
	description string
	currency    string
}

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello!")
}

func CreateMovement(c *fiber.Ctx) error {
	movement := new(models.Movement)
	if err := c.BodyParser(movement); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&movement)

	return c.Status(200).JSON(movement)
}
