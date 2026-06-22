package models

// RegisterRequest binds incoming registration JSON payloads.
type RegisterRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"` // 'User', 'Worker', etc.
}

// LoginRequest binds incoming login credential JSON payloads.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse returns a unified token schema to the client frontend.
type AuthResponse struct {
	UserID string `json:"user_id"` // Converted to string UUID for the frontend
	Token  string `json:"token"`
}
