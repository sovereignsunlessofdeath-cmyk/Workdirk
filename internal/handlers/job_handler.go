package handlers

import (
	"encoding/json"
	"net/http"
	models "workdirk/internal/model"
	"workdirk/internal/services"

	"github.com/google/uuid"
)

type JobHandler struct {
	jobSvc *services.JobService
}

func NewJobHandler(svc *services.JobService) *JobHandler {
	return &JobHandler{jobSvc: svc}
}

func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	custUUID, err := uuid.Parse(req.CustomerID)
	if err != nil {
		http.Error(w, "Invalid customer ID format", http.StatusBadRequest)
		return
	}
	workerUUID, err := uuid.Parse(req.WorkerID)
	if err != nil {
		http.Error(w, "Invalid worker ID format", http.StatusBadRequest)
		return
	}

	custBytes, _ := custUUID.MarshalBinary()
	workerBytes, _ := workerUUID.MarshalBinary()

	job, err := h.jobSvc.CreateNewJob(custBytes, workerBytes, req.Title, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
