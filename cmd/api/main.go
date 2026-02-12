package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yourusername/go_server/internal/config"
	"github.com/yourusername/go_server/internal/handlers"
	"github.com/yourusername/go_server/internal/middleware"
	"github.com/yourusername/go_server/internal/repository"
	"github.com/yourusername/go_server/internal/router"
	"github.com/yourusername/go_server/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize repository (database layer)
	repo := repository.NewRepository()

	// Initialize service layer
	svc := service.NewService(repo)

	// Initialize handlers
	handler := handlers.NewHandler(svc)

	// Setup router
	r := router.SetupRouter(handler, middleware.NewMiddleware())

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s in %s mode", port, cfg.Environment)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
