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
		return nil, fmt.Errorf("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return u, nil
}
