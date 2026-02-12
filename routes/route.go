package routes

import (
	"go_server/adapters"
	userRoutes "go_server/routes/user"

	"github.com/gofiber/fiber/v3"
)

// SetupRoutes configures all application routes (Base Router)
func SetupRoutes(app *fiber.App, orderHandler *adapters.HttpOrderHandler) {
	// Health check / Root route
	app.Get("/api", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Go Server API",
			"status":  "running",
			"version": "1.0.0",
		})
	})

	// Setup module-specific routes
	userRoutes.SetupUserRoutes(app, orderHandler)
}
