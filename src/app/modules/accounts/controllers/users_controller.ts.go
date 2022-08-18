package controllers

import (
	"base-fiber-api/src/app/modules/accounts/models"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Get(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Store(c *fiber.Ctx) error {
	user := models.User{}

	return c.JSON(user)
}

func Edit(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
