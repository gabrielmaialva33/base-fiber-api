package controllers

import (
	"base-fiber-api/src/app/modules/accounts/interfaces"
	"base-fiber-api/src/app/modules/accounts/models"
	"base-fiber-api/src/app/shared/pkg"
	"base-fiber-api/src/app/shared/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/imdario/mergo"
	"strconv"
	"strings"
)

type RoleServices struct {
	rr interfaces.RoleInterface
}

func RolesController(rr interfaces.RoleInterface) *RoleServices {
	return &RoleServices{rr}
}

func (r RoleServices) List(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage, _ := strconv.Atoi(c.Query("per_page", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")

	roles, err := r.rr.List(pkg.Meta{
		CurrentPage: page,
		PerPage:     perPage,
		Search:      search,
		Sort:        sort,
		Order:       order,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while getting roles",
			"error":   err.Error(),
		})
	}

	return c.JSON(roles)
}

func (r RoleServices) Get(c *fiber.Ctx) error {
	uuid := c.Params("id")

	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	role, err := r.rr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(role.PublicRole())
}

func (r RoleServices) Store(c *fiber.Ctx) error {
	data := models.Role{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing body",
			"error":   err.Error(),
		})
	}

	role := models.Role{
		Name: strings.ToLower(data.Slug),
	}
	if err := mergo.Merge(&role, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
		})
	}

	if errors := validators.ValidateStruct(role); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	newRole, err := r.rr.Store(&role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while creating role",
			"error":   err.Error(),
		})
	}

	return c.JSON(newRole.PublicRole())
}

func (r RoleServices) Edit(c *fiber.Ctx) error {
	uuid := c.Params("id")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	data := models.Role{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error while parsing data",
			"error":   err.Error(),
		})
	}

	role, err := r.rr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
			"error":   err.Error(),
		})
	}

	emptyRole := models.Role{
		Id:   role.Id,
		Name: strings.ToLower(data.Slug),
	}
	if err := mergo.Merge(&emptyRole, data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while merging data",
			"error":   err.Error(),
		})
	}

	if errors := validators.ValidatePartialStruct(emptyRole); len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	editedRole, err := r.rr.Edit(&emptyRole)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while updating user",
			"error":   err.Error(),
		})
	}

	return c.JSON(editedRole.PublicRole())
}

func (r RoleServices) Delete(c *fiber.Ctx) error {
	uuid := c.Params("id")
	if validators.ValidateUUID(uuid) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid UUID",
		})
	}

	role, err := r.rr.Get(uuid)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	deleteRole := models.Role{
		Id:   role.Id,
		Name: "deleted:" + role.Name + ":" + strings.Split(role.Id, "-")[0],
		Slug: "deleted:" + role.Slug + ":" + strings.Split(role.Id, "-")[0],
	}

	if err := r.rr.Delete(&deleteRole); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while deleting user",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Role deleted",
	})
}
