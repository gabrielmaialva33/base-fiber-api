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
	api.Get("/:id", controller.Get).Name("roles.get")
	api.Post("/", controller.Store).Name("roles.store")
	api.Put("/:id", controller.Edit).Name("roles.edit")
	api.Delete("/:id", controller.Delete).Name("roles.delete")
}
