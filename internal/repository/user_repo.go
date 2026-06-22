package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID           []byte // BINARY(16)
	Name         string
	Email        string
	PasswordHash []byte // BINARY(60)
	PhoneNumber  string
	Role         string
}

type UserRepository struct {
	dbConn *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{dbConn: db}
}

func (r *UserRepository) Create(name, email string, passwordHash []byte, phone, role string) (*User, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed generating binary uuid: %w", err)
	}
	binaryID, _ := newUUID.MarshalBinary()

	query := `INSERT INTO users (id, name, email, password_hash, phone_number, role) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = r.dbConn.Exec(query, binaryID, name, email, passwordHash, phone, role)
	if err != nil {
		return nil, fmt.Errorf("failed executing mysql insert: %w", err)
	}

	return &User{ID: binaryID, Name: name, Email: email, PasswordHash: passwordHash, PhoneNumber: phone, Role: role}, nil
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, password_hash, phone_number, role FROM users WHERE email = ? LIMIT 1`
	var u User
	err := r.dbConn.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.PhoneNumber, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
