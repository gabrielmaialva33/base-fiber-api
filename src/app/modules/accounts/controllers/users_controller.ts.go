package controllers

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/utils"
	"github.com/gofiber/fiber/v2"
)

type UserServices struct {
	ur interfaces.UserInterface
}

func UsersController(ur interfaces.UserInterface) *UserServices {
	return &UserServices{ur}
}

func List(c *fiber.Ctx) error {
	return c.SendString("List users!")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Get user!")
}

func (s *UserServices) Store(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user := models.User{
		FirstName:       data["first_name"],
		LastName:        data["last_name"],
		Email:           data["email"],
		UserName:        data["user_name"],
		Password:        data["password"],
		ConfirmPassword: data["confirm_password"],
	}
	if errors := utils.ValidateStruct(user); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	newUser, err := s.ur.Store(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while creating user",
			"error":   err,
		})
	}

	return c.JSON(newUser)
}

func Edit(c *fiber.Ctx) error {
	return c.SendString("Edit user!")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Delete user!")
}
