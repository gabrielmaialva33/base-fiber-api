package controllers

import (
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/utils"
	"base-fiber-api/src/database"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	return c.SendString("List users!")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Get user!")
}

func Store(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Edit(c *fiber.Ctx) error {
	return c.SendString("Edit user!")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Delete user!")
}
