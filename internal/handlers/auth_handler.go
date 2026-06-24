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

// 1. LOGIN METHOD
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

	// This sends the data directly over to your service layer structure
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
		UserID: string(user.ID),
		Token:  session.Token,
	})
}

// 2. LOGOUT METHOD
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract token from the Authorization header (e.g., "Bearer <token>")
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		http.Error(w, "Missing authorization token", http.StatusUnauthorized)
		return
	}

	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}

	// Revoke session in the database
	err := h.sessionSvc.RevokeSession(tokenStr)
	if err != nil {
		http.Error(w, "Failed to terminate session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully logged out"}`))
}

// 3. FORGOT PASSWORD METHOD
func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// We expect a JSON payload containing only {"email": "user@example.com"}
	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		http.Error(w, "Email address is required", http.StatusBadRequest)
		return
	}

	// Trigger your service to find user, generate token, and send email link
	// Note: You will need to define this GeneratePasswordResetToken method in your AuthService!
	err := h.authSvc.GeneratePasswordResetToken(req.Email)
	if err != nil {
		// Security tip: You can choose to return an error or hide it so attackers don't scan emails
		http.Error(w, "Failed to initiate password reset process", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Password reset instructions have been sent to your email"}`))
}
