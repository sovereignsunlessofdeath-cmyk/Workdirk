package handlers

import (
	"encoding/json"
	"net/http"
	models "workdirk/internal/model"
	"workdirk/internal/services"

	"github.com/google/uuid"
)

type ReviewHandler struct {
	reviewSvc *services.ReviewService
}

func NewReviewHandler(svc *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewSvc: svc}
}

func (h *ReviewHandler) LeaveReview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.SubmitReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	jobUUID, _ := uuid.Parse(req.JobID)
	revUUID, _ := uuid.Parse(req.ReviewerID)
	reveeUUID, _ := uuid.Parse(req.RevieweeID)

	jobB, _ := jobUUID.MarshalBinary()
	revB, _ := revUUID.MarshalBinary()
	reveeB, _ := reveeUUID.MarshalBinary()

	review, err := h.reviewSvc.SubmitReview(jobB, revB, reveeB, req.Rating, req.Comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}
