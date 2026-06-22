package handlers

import (
	"encoding/json"
	"net/http"
	models "workdirk/internal/model"
	"workdirk/internal/services"
)

type SessionHandler struct {
	sessionSvc *services.SessionService
}

func NewSessionHandler(svc *services.SessionService) *SessionHandler {
	return &SessionHandler{sessionSvc: svc}
}

// CheckSession acts as an endpoint placeholder to verify incoming authorization states.
func (h *SessionHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.SessionVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Session check routes can be expanded here as token validation logic grows.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "active"}`))
}
