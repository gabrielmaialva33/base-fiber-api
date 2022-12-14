package middlewares

import (
	"base-fiber-api/src/app/shared/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Auth(c *fiber.Ctx) error {
	bear := c.Get("Authorization")
	if bear == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	token := strings.Split(bear, " ")[1]
	if _, err := utils.ParseJWT(token); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
