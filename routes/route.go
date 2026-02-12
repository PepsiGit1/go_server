package routes

import (
	"go_server/adapters"

	"github.com/gofiber/fiber/v3"
)

// SetupRoutes configures all application routes
func SetupRoutes(app *fiber.App, orderHandler *adapters.HttpOrderHandler) {
	// Order routes
	app.Post("/order", orderHandler.CreateOrder)
}
