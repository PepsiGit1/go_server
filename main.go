package main

import (
	"fmt"
	"log"
	"os"

	"go_server/adapters"
	"go_server/config"
	core "go_server/core/order"
	"go_server/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	app := fiber.New()

	// Build PostgreSQL DSN from environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("PG_HOST", "localhost"),
		getEnv("PG_USER", "postgres"),
		getEnv("PG_PASSWORD", "1234"),
		getEnv("PG_DB", "postgres"),
		getEnv("PG_PORT", "5432"),
	)

	// Initialize GORM with PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL database!")

	// Auto-migrate the schema
	if err := db.AutoMigrate(&core.Order{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Store db in config for potential reuse
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	config.DB = sqlDB

	// Set up the core service and adapters
	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	// Setup routes
	routes.SetupRoutes(app, orderHandler)

	// Start the server
	port := getEnv("PORT", "8000")
	log.Printf("Server starting on port %s...", port)
	app.Listen(":" + port)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
