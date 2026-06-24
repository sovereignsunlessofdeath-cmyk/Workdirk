package models

// UserProfileResponse structures clean user data for account or profile screens.
type UserProfileResponse struct {
	ID           string `json:"id"` // Converted from BINARY(16) bytes to string UUID
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Role         string `json:"role"`
	PasswordHash string `json:"-"` // Exclude from JSON responses for security
}

// UpdateProfileRequest binds incoming JSON payloads for updating user account details.
type UpdateProfileRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
