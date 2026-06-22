package handlers

import (
	"encoding/json"
	"net/http"
	models "workdirk/internal/model"
	"workdirk/internal/services"
)

type AuthHandler struct {
	authSvc    *services.AuthService
	sessionSvc *services.SessionService
}

func NewAuthHandler(authSvc *services.AuthService, sessionSvc *services.SessionService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc, sessionSvc: sessionSvc}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.authSvc.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	session, err := h.sessionSvc.GenerateSession(user.ID)
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.AuthResponse{
		UserID: string(user.ID), // In a real client pipeline, you can use uuid.FromBytes to make it a string
		Token:  session.Token,
	})
}
