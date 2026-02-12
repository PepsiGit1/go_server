package router

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/go_server/internal/handlers"
	"github.com/yourusername/go_server/internal/middleware"
)

func SetupRouter(h *handlers.Handler, m *middleware.Middleware) *mux.Router {
	r := mux.NewRouter()

	// Apply global middleware
	r.Use(m.Logger)
	r.Use(m.CORS)

	// API v1 routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// Health check
	api.HandleFunc("/health", h.HealthCheck).Methods("GET")

	// User routes
	users := api.PathPrefix("/users").Subrouter()
	users.HandleFunc("", h.GetUsers).Methods("GET")
	users.HandleFunc("", h.CreateUser).Methods("POST")
	users.HandleFunc("/{id}", h.GetUser).Methods("GET")
	users.HandleFunc("/{id}", h.UpdateUser).Methods("PUT")
	users.HandleFunc("/{id}", h.DeleteUser).Methods("DELETE")

	// Protected routes (require authentication)
	protected := api.PathPrefix("/protected").Subrouter()
	protected.Use(m.Auth)
	protected.HandleFunc("/profile", h.GetProfile).Methods("GET")

	return r
}
