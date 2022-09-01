package controllers

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
	"base-fiber-api/src/app/shared/utils"
	"base-fiber-api/src/app/shared/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"strconv"
	"strings"
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

	return c.JSON(user.PublicUser())
}

func (s *UserServices) Store(c *fiber.Ctx) error {
	data := models.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
		})
	}

	user := models.User{}
	if err := mergo.Merge(&user, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
		})
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

	return c.JSON(newUser.PublicUser())
}

func (s *UserServices) Edit(c *fiber.Ctx) error {
	uuid := c.Params("id")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	data := models.User{}
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

	emptyUser := models.User{}
	if err := mergo.Merge(&emptyUser, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
		})
	}

	if errors := validators.ValidatePartialStruct(emptyUser); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	editedUser, err := s.ur.Edit(user.Id, &emptyUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while updating user",
			"error":   err,
		})
	}

	return c.JSON(editedUser.PublicUser())
}

func (s *UserServices) Delete(c *fiber.Ctx) error {
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

	deleteUser := models.User{
		Email:    "deleted:" + user.Email + ":" + strings.Split(user.Id, "-")[0],
		UserName: "deleted:" + user.UserName + ":" + strings.Split(user.Id, "-")[0],
	}

	if err := s.ur.Delete(user.Id, &deleteUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while deleting user",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted",
	})
}

func (s *UserServices) Login(c *fiber.Ctx) error {
	data := map[string]string{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
		})
	}

	user, err := s.ur.FindBy("user_name", data["uid"])
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	match, err := pkg.ComparePasswordAndHash(data["password"], user.Password)
	if err != nil || match == false {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while comparing password",
		})
	}

	token, err := utils.GenerateJwt(user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while generating token",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
