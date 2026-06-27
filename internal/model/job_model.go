package model

// CreateJobRequest captures details required to spin up a work milestone.
type CreateJobRequest struct {
	CustomerID  string `json:"customer_id"` // Input as a string UUID
	WorkerID    string `json:"worker_id"`   // Input as a string UUID
	Title       string `json:"title"`
	Description string `json:"description"`
}

// JobResponse structures clean job outputs without raw binary bytes.
type JobResponse struct {
	ID          string `json:"id"`
	CustomerID  string `json:"customer_id"`
	WorkerID    string `json:"worker_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
