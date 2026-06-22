package models

// InitiateEscrowRequest binds incoming project funding commands.
type InitiateEscrowRequest struct {
	JobID      string  `json:"job_id"`
	CustomerID string  `json:"customer_id"`
	WorkerID   string  `json:"worker_id"`
	Amount     float64 `json:"amount"`
}

// PaymentResponse standardizes outgoing escrow balance responses.
type PaymentResponse struct {
	ID         string  `json:"id"`
	JobID      string  `json:"job_id"`
	CustomerID string  `json:"customer_id"`
	WorkerID   string  `json:"worker_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}
