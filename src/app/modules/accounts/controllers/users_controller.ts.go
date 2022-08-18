package controllers

import (
	"base-fiber-api/src/app/modules/accounts/models"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	return c.SendString("List users!")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Get user!")
}

func Store(c *fiber.Ctx) error {
	user := models.User{}

	return c.JSON(user)
}

func Edit(c *fiber.Ctx) error {
	return c.SendString("Edit user!")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Delete user!")
}
