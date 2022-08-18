package main

import (
	"base-fiber-api/src/app/modules/accounts/routes"
	"base-fiber-api/src/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
	})

	app.Use(logger.New())

	database.Connect()

	routes.UserRoutes(app)

	_ = app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
