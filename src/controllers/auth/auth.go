package auth_controllers

import (
	"encoding/json"
	auth_model "smol-api/src/models/auth"
	user_model "smol-api/src/models/users"
	jwt_services "smol-api/src/services"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthUser(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	var authUser user_model.User
	var reqUser auth_model.AuthRequest

	if err := c.BodyParser(&reqUser); err != nil {
		return err
	}

	if err := db.Where("email = ?", reqUser.Email).First(&authUser).Error; err != nil {
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(reqUser.Password))
	if err != nil {
		return err
	}

	token, err := jwt_services.GenerateJWT(authUser.ID, authUser.Role)
	if err != nil {
		return err
	}

	data := struct {
		Token string          `json:"token"`
		User  user_model.User `json:"user"`
	}{
		Token: token,
		User:  authUser,
	}

	userJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).Send(userJSON)
}
