package handlers

import (
	"encoding/json"
	"net/http"
	models "workdirk/internal/model"
	"workdirk/internal/services"

	"github.com/google/uuid"
)

type PaymentHandler struct {
	paymentSvc *services.PaymentService
}

func NewPaymentHandler(svc *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentSvc: svc}
}

func (h *PaymentHandler) InitializeEscrow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.InitiateEscrowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	jobUUID, _ := uuid.Parse(req.JobID)
	custUUID, _ := uuid.Parse(req.CustomerID)
	workUUID, _ := uuid.Parse(req.WorkerID)

	jobB, _ := jobUUID.MarshalBinary()
	custB, _ := custUUID.MarshalBinary()
	workB, _ := workUUID.MarshalBinary()

	payment, err := h.paymentSvc.HoldFunds(jobB, custB, workB, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}
