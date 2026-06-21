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

// 💡 ADDED: GetByEmail method now belongs directly to AuthRepository
func (r *AuthRepository) GetByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, password_hash, phone_number, role FROM users WHERE email = ? LIMIT 1`
	var u User
	err := r.dbConn.QueryRow(query, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.PhoneNumber, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
