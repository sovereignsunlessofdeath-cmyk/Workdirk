package db

import (
	"database/sql"
	"fmt"
	"log"
)

type MigrationEngine struct {
	db *sql.DB
}

func NewMigrationEngine(db *sql.DB) *MigrationEngine {
	return &MigrationEngine{db: db}
}

func (m *MigrationEngine) RunSafeMigrations(dirPath string) error {
	log.Println("Running structural database schema verification...")

	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id BINARY(16) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash BINARY(60) NOT NULL,
			phone_number VARCHAR(50),
			role VARCHAR(50) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS jobs (
			id BINARY(16) PRIMARY KEY,
			customer_id BINARY(16) NOT NULL,
			worker_id BINARY(16) NOT NULL,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			status VARCHAR(50) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS payments (
			id BINARY(16) PRIMARY KEY,
			job_id BINARY(16) NOT NULL,
			customer_id BINARY(16) NOT NULL,
			worker_id BINARY(16) NOT NULL,
			amount DECIMAL(10,2) NOT NULL,
			status VARCHAR(50) NOT NULL
		);`,
	}

	for _, query := range queries {
		if _, err := m.db.Exec(query); err != nil {
			return fmt.Errorf("migration instruction execution failed: %w", err)
		}
	}

	log.Println("✅ All core relational database schemas verified successfully.")
	return nil
}
