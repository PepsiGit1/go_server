package user

import (
	"go_server/adapters"

	"github.com/gofiber/fiber/v3"
)

// SetupUserRoutes configures all user-related routes
func SetupUserRoutes(app *fiber.App, userHandler *adapters.HttpOrderHandler) {
	// User routes
	app.Get("/users", userHandler.GetAllUsers)
}
