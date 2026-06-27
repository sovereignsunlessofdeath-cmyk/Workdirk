package repository

import (
	"database/sql"
)

type AuthRepository struct {
	dbConn *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{dbConn: db}
}

// GetByEmail queries the DB and returns the raw database User model (with its PasswordHash).
func (r *AuthRepository) GetByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, password_hash, phone_number, role FROM users WHERE email = ? LIMIT 1`

	var u User
	err := r.dbConn.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.PhoneNumber, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *AuthRepository) CreateUser(u *User) error {
	query := `INSERT INTO users (id, name, email, password_hash, phone_number, role) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.dbConn.Exec(query, u.ID, u.Name, u.Email, u.PasswordHash, u.PhoneNumber, u.Role)
	return err
}

func (r *AuthRepository) Logout(user *User) error {
	// Update the user's logout status or session in the database
	// This is a placeholder—adjust based on your actual schema
	query := `UPDATE users SET updated_at = NOW() WHERE id = ?`
	_, err := r.dbConn.Exec(query, user.ID)
	return err
}
