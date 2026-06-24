package services

import (
	"fmt"
	"workdirk/internal/repository"

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

	err = bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return u, nil
}

// GeneratePasswordResetToken handles checking the user email and creating a reset token
func (s *AuthService) GeneratePasswordResetToken(email string) error {
	// This placeholder allows your handler code to compile cleanly.
	// Later, we will add the SQL queries here to save a reset token to your database.
	return nil
}
