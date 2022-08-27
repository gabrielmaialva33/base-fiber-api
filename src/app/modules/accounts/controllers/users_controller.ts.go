package controllers

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
	"base-fiber-api/src/app/shared/validators"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserServices struct {
	ur interfaces.UserInterface
}

func UsersController(ur interfaces.UserInterface) *UserServices {
	return &UserServices{ur}
}

func (s *UserServices) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	order := c.Query("order", "id")

	users, err := s.ur.List(pkg.Pagination{
		Page:    page,
		PerPage: perPage,
		Order:   order,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while getting users",
			"error":   err,
		})
	}

	return c.JSON(users)
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
	if errors := validators.ValidateStruct(user); len(errors) > 0 {
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
