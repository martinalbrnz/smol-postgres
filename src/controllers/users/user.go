package user_controllers

import (
	"encoding/json"
	"fmt"
	user_model "smol-api/src/models/users"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var users []user_model.User
	db.Find(&users)

	for _, user := range users {
		fmt.Printf("%s\n", user.Email)
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).Send(usersJSON)
}

func GetUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var user user_model.User
	if err := db.Where("id = ?", c.Params("id")).First(&user).Error; err != nil {
		return err
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).Send(userJSON)
}

func CreateUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var user user_model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).Send(userJSON)
}
