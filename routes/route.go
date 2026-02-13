package routes

import (
	"go_server/adapters"
	userRoutes "go_server/routes/user"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, userHandler *adapters.HttpUserHandler) {
	app.Get("/api", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Go Server API",
			"status":  "running",
			"version": "1.0.0",
		})
	})

	// Setup module-specific routes
	userRoutes.SetupUserRoutes(app, userHandler)
}
