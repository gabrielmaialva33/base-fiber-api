package routes

import (
	"base-fiber-api/src/app/modules/accounts/http/controllers"
	"base-fiber-api/src/app/shared/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, controller *controllers.UserServices) {
	app.Post("/login", controller.Login).Name("users.login")

	api := app.Group("/users")
	app.Post("/", controller.Store).Name("users.store").Name("users.store")

	api.Use(middlewares.Auth)

	api.Get("/", controller.List).Name("users.list")
	api.Get("/:id", controller.Get).Name("users.get")
	api.Put("/:id", controller.Edit).Name("users.edit")
	api.Delete("/:id", controller.Delete).Name("users.delete")
}
