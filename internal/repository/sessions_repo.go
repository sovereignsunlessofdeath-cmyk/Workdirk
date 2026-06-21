package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        []byte // BINARY(16)
	UserID    []byte // BINARY(16)
	Token     string
	ExpiresAt time.Time
}

type SessionRepository struct {
	dbConn *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{dbConn: db}
}

func (r *SessionRepository) CreateSession(userID []byte, token string, expiresAt time.Time) (*Session, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	binaryID, _ := newUUID.MarshalBinary()

	query := `INSERT INTO sessions (id, user_id, token, expires_at) VALUES (?, ?, ?, ?)`
	_, err = r.dbConn.Exec(query, binaryID, userID, token, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &Session{ID: binaryID, UserID: userID, Token: token, ExpiresAt: expiresAt}, nil
}
