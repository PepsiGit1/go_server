package user

import (
	"go_server/adapters"

	"github.com/gofiber/fiber/v3"
)

// SetupUserRoutes configures all user-related routes
func SetupUserRoutes(app *fiber.App, userHandler *adapters.HttpUserHandler) {
	// User routes
	app.Get("/users", userHandler.GetAllUsers)
	app.Post("/register", userHandler.Register)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Post("/login", userHandler.LoginUser)

}
