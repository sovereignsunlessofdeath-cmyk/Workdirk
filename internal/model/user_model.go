package model

// 1. Core Database Entity Model (Add this!)
type User struct {
    ID       int    `json:"id" db:"id"`
    Username string `json:"username" db:"username"`
    Email    string `json:"email" db:"email"`
    Password string `json:"password" db:"password"`
    Phone    string `json:"???" db:"phone"` // <-- We need to see this exact JSON tag!
}

// UserProfileResponse structures clean user data for account or profile screens.
type UserProfileResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	// Note: We removed PasswordHash completely from here since the database entity 'User' handles it!
}

// UpdateProfileRequest binds incoming JSON payloads for updating user account details.
type UpdateProfileRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
