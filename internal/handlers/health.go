package handlers

import (
	"net/http"
)

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{
		"status":  "healthy",
		"service": "go-server-api",
	})
}
