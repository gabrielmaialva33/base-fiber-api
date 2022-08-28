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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting users",
			"error":   err.Error(),
		})
	}

	return c.JSON(users)
}

func (s *UserServices) Get(c *fiber.Ctx) error {
	uuid := c.Params("id")

	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	user, err := s.ur.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user)
}

func (s *UserServices) Store(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
		})
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

func (s *UserServices) Edit(c *fiber.Ctx) error {
	uuid := c.Params("id")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
		})
	}

	user, err := s.ur.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	user.FirstName = data["first_name"]
	user.LastName = data["last_name"]
	user.Email = data["email"]
	user.UserName = data["user_name"]
	user.Password = data["password"]
	user.ConfirmPassword = data["confirm_password"]

	if errors := validators.ValidateStruct(user); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	updatedUser, err := s.ur.Edit(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while updating user",
			"error":   err,
		})
	}

	return c.JSON(updatedUser)
}

func (s *UserServices) Delete(c *fiber.Ctx) error {
	uuid := c.Params("id")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	return c.JSON(uuid)
}
