package middlewares

import (
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/utils"
	"base-fiber-api/src/database"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Acl(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bear := c.Get("Authorization")
		if bear == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token := strings.Split(bear, " ")[1]
		id, err := utils.ParseJWT(token)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		var user models.User
		if err := database.DB.Preload("Roles").Where("id = ?", id).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		for _, role := range user.Roles {
			if !utils.Contains(allowedRoles, strings.ToLower(role.Name)) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Unauthorized",
				})
			}
		}

		return c.Next()
	}
}
