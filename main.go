package main

import (
	"log"
	"os"

	"go_server/adapters"
	"go_server/config"
	core "go_server/core/user"
	"go_server/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Initialize database connection using config
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.CloseDB()

	// Auto-migrate the schema
	if err := db.AutoMigrate(&core.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Set up the core service and adapters
	userRepo := adapters.NewGormUserRepository(db)
	orderService := core.NewUserService(userRepo)
	orderHandler := adapters.NewHttpUserHandler(orderService)

	// Setup routes
	routes.SetupRoutes(app, orderHandler)

	// Start the server
	port := getEnv("PORT", "8000")
	log.Printf("Server starting on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
