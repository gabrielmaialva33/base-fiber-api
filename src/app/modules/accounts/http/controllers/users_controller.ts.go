package controllers

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/modules/accounts/services"
	"base-fiber-api/src/app/shared/pkg"
	"base-fiber-api/src/app/shared/utils"
	"base-fiber-api/src/app/shared/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"strconv"
	"strings"
)

// UsersController is the controller for users
type UsersController struct {
	ur services.UserServicesInterface
}

// NewUsersController creates a new instance of the user controller
func NewUsersController(ur interfaces.UserInterface) *UsersController {
	return &UsersController{ur}
}

func (u *UsersController) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")

	users, err := u.ur.List(pkg.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Search:      search,
		Sort:        sort,
		Order:       order,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting users",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": true,
		})
	}

	return c.JSON(users)
}

func (u *UsersController) Get(c *fiber.Ctx) error {
	uuid := c.Params("userId")

	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	user, err := u.ur.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(user.PublicUser())
}

func (u *UsersController) Store(c *fiber.Ctx) error {
	data := models.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
		})
	}

	user := models.User{
		Role: models.RoleUser,
	}
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

	newUser, err := u.ur.Store(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while creating user",
			"error":   err.Error(),
		})
	}

	return c.JSON(newUser.PublicUser())
}

func (u *UsersController) Edit(c *fiber.Ctx) error {
	uuid := c.Params("userId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	data := models.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
		})
	}

	user, err := u.ur.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	emptyUser := models.User{
		Id: user.Id,
	}
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

	editedUser, err := u.ur.Edit(&emptyUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while updating user",
			"error":   err.Error(),
		})
	}

	return c.JSON(editedUser.PublicUser())
}

func (u *UsersController) Delete(c *fiber.Ctx) error {
	uuid := c.Params("userId")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	user, err := u.ur.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	deleteUser := models.User{
		Id:       user.Id,
		Email:    "deleted:" + user.Email + ":" + strings.Split(user.Id, "-")[0],
		UserName: "deleted:" + user.UserName + ":" + strings.Split(user.Id, "-")[0],
	}

	if err := u.ur.Delete(&deleteUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while deleting user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "User deleted",
	})
}

func (u *UsersController) SignIn(c *fiber.Ctx) error {
	data := models.Login{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(data); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	user, err := u.ur.FindManyBy([]string{"email", "user_name"}, data.Uid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"status":  fiber.StatusNotFound,
			"display": true,
			"error":   err.Error(),
		})
	}

	if match, _ := pkg.ComparePasswordAndHash(data.Password, user.Password); match == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
			"status":  fiber.StatusUnauthorized,
			"display": true,
		})
	}

	token, err := utils.GenerateJwt(user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while generating token",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"status":  fiber.StatusOK,
		"display": false,
		"user":    user.PublicUser(),
		"token":   token,
	})
}

func (u *UsersController) SignUp(c *fiber.Ctx) error {
	data := models.User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
			"status":  fiber.StatusBadRequest,
			"display": false,
		})
	}

	user := models.User{}
	if err := mergo.Merge(&user, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	if errors := validators.ValidateStruct(user); len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
			"status":  fiber.StatusUnprocessableEntity,
			"display": true,
		})
	}

	newUser, err := u.ur.Store(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while creating user",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	token, err := utils.GenerateJwt(newUser.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while generating token",
			"error":   err.Error(),
			"status":  fiber.StatusInternalServerError,
			"display": false,
		})
	}

	return c.JSON(fiber.Map{
		"message": "SignUp successful. Please login to continue",
		"status":  fiber.StatusOK,
		"display": false,
		"user":    newUser.PublicUser(),
		"token":   token,
	})
}
