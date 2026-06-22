package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Payment struct {
	ID         []byte // BINARY(16)
	JobID      []byte // BINARY(16)
	CustomerID []byte // BINARY(16)
	WorkerID   []byte // BINARY(16)
	Amount     float64
	Status     string
}

type PaymentRepository struct {
	dbConn *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{dbConn: db}
}

func (r *PaymentRepository) CreatePayment(jobID, customerID, workerID []byte, amount float64) (*Payment, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	binaryID, _ := newUUID.MarshalBinary()
	defaultStatus := "Held_Escrow"

	query := `INSERT INTO payments (id, job_id, customer_id, worker_id, amount, status) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = r.dbConn.Exec(query, binaryID, jobID, customerID, workerID, amount, defaultStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return &Payment{ID: binaryID, JobID: jobID, CustomerID: customerID, WorkerID: workerID, Amount: amount, Status: defaultStatus}, nil
}

func (r *PaymentRepository) UpdateStatus(paymentID []byte, newStatus string) error {
	query := `UPDATE payments SET status = ? WHERE id = ?`
	_, err := r.dbConn.Exec(query, newStatus, paymentID)
	return err
}
