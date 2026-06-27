package services

import (
	"fmt"
	"workdirk/internal/repository"

	model "workdirk/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo *repository.AuthRepository
}

// 💡 FIX: Change the input argument type to *repository.AuthRepository here:
func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{authRepo: repo}
}

func (s *AuthService) Login(email, password string) (*repository.User, error) {
	u, err := s.authRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email ")
	}

	// 1. Run the standard comparison
	// ❌ Old version with quotes causing the error:
	// bcryptErr := bcrypt.CompareHashAndPassword(u."_", []byte(password))

	// ✅ Correct version (no quotes around the field name):
	bcryptErr := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	// 2. 👇 CALL THE ONE-LINE DEBUG HELPER HERE 👇
	logAuthDebug("Service.Login Check", password, string(u.PasswordHash), bcryptErr)

	if bcryptErr != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return u, nil
}

// 👇 Make sure this helper function sits at the very bottom of the file
func logAuthDebug(location, inputPassword, databaseHash string, bcryptError error) {
	status := "✅ MATCH"
	if bcryptError != nil {
		status = fmt.Sprintf("❌ MISMATCH (%v)", bcryptError)
	}

	// Prints everything on a single, easy-to-read line with timestamps
	fmt.Printf("[AUTH] %s | Input: %q | DB Hash: %q | Result: %s", location, inputPassword, databaseHash, status)
}

func (s *AuthService) Register(name, email, password, phoneNumber string) (*model.UserProfileResponse, error) {
	// 1. Hash the plain-text password securely
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to process password: %w", err)
	}

	// 2. Build the user model entity structure
	newUser := &model.UserProfileResponse{
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        "user", // Assign a default role
	}

	// 3. Pass the completed model down to your repository layer to save it
	// Note: Ensure your repository has a matching 'Create' or 'Save' method!
	err = s.authRepo.CreateUser(&repository.User{
		Name:         newUser.Name,
		Email:        newUser.Email,
		PasswordHash: hashedPassword, // Store the hashed password (as []byte)
		PhoneNumber:  newUser.PhoneNumber,
		Role:         newUser.Role,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create user account: %w", err)
	}

	return newUser, nil
}

func (s *AuthService) Logout(user *repository.User) error {
	// Execute any business validation rule if needed before logout
	if user == nil {
		return fmt.Errorf("cannot log out an unauthenticated user")
	}

	// Pass the model down to your repository to flip their status column in MySQL
	err := s.authRepo.Logout(user)
	if err != nil {
		return fmt.Errorf("failed to update logout status in database: %w", err)
	}

	return nil
}

// GeneratePasswordResetToken handles checking the user email and creating a reset token
func (s *AuthService) GeneratePasswordResetToken(email string) error {
	// This placeholder allows your handler code to compile cleanly.
	// Later, we will add the SQL queries here to save a reset token to your database.
	return nil
}
