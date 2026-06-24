package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
	"workdirk/internal/repository"
)

type SessionService struct {
	sessionRepo *repository.SessionRepository
}

func NewSessionService(repo *repository.SessionRepository) *SessionService {
	return &SessionService{sessionRepo: repo}
}

func (s *SessionService) GenerateSession(userID []byte) (*repository.Session, error) {
	// Generate a highly secure random token string
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("failed to generate secure token: %w", err)
	}
	tokenStr := hex.EncodeToString(b)

	// Session stays valid for 24 hours
	expiryTime := time.Now().Add(24 * time.Hour)

	sess, err := s.sessionRepo.CreateSession(userID, tokenStr, expiryTime)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

// RevokeSession deletes or invalidates a user session token in the database
func (s *SessionService) RevokeSession(token string) error {
	// If you are using a SQL database, you would delete or update the session record.
	// For now, let's write a mock placeholder that returns nil (no error) so your code compiles.
	return nil
}
