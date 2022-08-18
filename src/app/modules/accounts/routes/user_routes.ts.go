package routes

import (
	"base-fiber-api/src/app/modules/accounts/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/users")

	api.Get("/", controllers.List).Name("users.list")
	api.Get("/:id", controllers.Get).Name("users.get")
	api.Post("/", controllers.Store).Name("users.store")
	api.Put("/:id", controllers.Edit).Name("users.edit")
	api.Delete("/:id", controllers.Delete).Name("users.delete")
}
