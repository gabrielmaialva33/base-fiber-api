package main

import (
	"base-fiber-api/src/app/modules/accounts/http/controllers"
	"base-fiber-api/src/app/modules/accounts/http/routes"
	"base-fiber-api/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	services := database.NewRepositories(dsn)
	services.Drop()
	services.Migrate()
	services.Seed()

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	app.Use(logger.New())

	userController := controllers.UsersController(services.User)
	roleController := controllers.RolesController(services.Role)
	routes.UserRoutes(app, userController)
	routes.RoleRoutes(app, roleController)

	_ = app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
