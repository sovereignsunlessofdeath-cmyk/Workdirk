package model

// SubmitReviewRequest captures incoming worker evaluations.
type SubmitReviewRequest struct {
	JobID      string `json:"job_id"`
	ReviewerID string `json:"reviewer_id"`
	RevieweeID string `json:"reviewee_id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
}

// ReviewResponse outputs structured feedback cards cleanly.
type ReviewResponse struct {
	ID         string `json:"id"`
	JobID      string `json:"job_id"`
	ReviewerID string `json:"reviewer_id"`
	RevieweeID string `json:"reviewee_id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
}
