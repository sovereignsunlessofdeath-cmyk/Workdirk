package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Job struct {
	ID          []byte // BINARY(16)
	CustomerID  []byte // BINARY(16)
	WorkerID    []byte // BINARY(16)
	Title       string
	Description string
	Status      string
}

type JobRepository struct {
	dbConn *sql.DB
}

func NewJobRepository(db *sql.DB) *JobRepository {
	return &JobRepository{dbConn: db}
}

func (r *JobRepository) CreateJob(customerID, workerID []byte, title, description string) (*Job, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	binaryID, _ := newUUID.MarshalBinary()
	defaultStatus := "Pending"

	query := `INSERT INTO jobs (id, customer_id, worker_id, title, description, status) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = r.dbConn.Exec(query, binaryID, customerID, workerID, title, description, defaultStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to create job: %w", err)
	}

	return &Job{ID: binaryID, CustomerID: customerID, WorkerID: workerID, Title: title, Description: description, Status: defaultStatus}, nil
}
