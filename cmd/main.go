package main

import (
	"base-fiber-api/src/app/modules/accounts/http/controllers"
	"base-fiber-api/src/app/modules/accounts/http/routes"
	"base-fiber-api/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	services := database.Connect(dsn)

	services.Drop()
	services.Migrate()
	services.Seed()

	app := fiber.New(fiber.Config{
		AppName:                 "Base Fiber API",
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Request-With",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	// Controllers
	userController := controllers.NewUsersController(services.User)
	roleController := controllers.NewRolesController(services.Role)

	// Routes
	routes.UserRoutes(app, userController)
	routes.RoleRoutes(app, roleController)

	_ = app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
