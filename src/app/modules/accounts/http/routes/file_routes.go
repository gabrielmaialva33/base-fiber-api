package routes

import (
	"base-fiber-api/src/app/modules/accounts/http/controllers"
	"base-fiber-api/src/app/shared/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func FileRoutes(app *fiber.App) {
	api := app.Group("/files")

	api.Use(middlewares.Acl([]string{"admin", "root", "user"}))

	api.Post("/", controllers.Store).Name("files.store")
	api.Static("/uploads", "public/uploads/")
}
