package routes

import (
	"base-fiber-api/src/app/modules/accounts/http/controllers"
	"base-fiber-api/src/app/shared/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RoleRoutes(app *fiber.App, controller *controllers.RoleServices) {
	api := app.Group("/roles")

	api.Use(middlewares.Auth)

	api.Get("/", controller.List).Name("roles.list")
	api.Get("/:roleId", controller.Get).Name("roles.get")
	api.Post("/", controller.Store).Name("roles.store")
	api.Put("/:roleId", controller.Edit).Name("roles.edit")
	api.Delete("/:roleId", controller.Delete).Name("roles.delete")
}
