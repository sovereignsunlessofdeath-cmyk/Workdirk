package services

import (
	"fmt"
	"workdirk/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) RegisterUser(name, email, password, phone, role string) (*repository.User, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	u, err := s.userRepo.Create(name, email, hashedBytes, phone, role)
	if err != nil {
		return nil, err
	}

	return u, nil
}
