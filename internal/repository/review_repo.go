package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Review struct {
	ID         []byte // BINARY(16)
	JobID      []byte // BINARY(16)
	ReviewerID []byte // BINARY(16)
	RevieweeID []byte // BINARY(16)
	Rating     int
	Comment    string
}

type ReviewRepository struct {
	dbConn *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{dbConn: db}
}

func (r *ReviewRepository) CreateReview(jobID, reviewerID, revieweeID []byte, rating int, comment string) (*Review, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	binaryID, _ := newUUID.MarshalBinary()

	query := `INSERT INTO reviews (id, job_id, reviewer_id, reviewee_id, rating, comment) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = r.dbConn.Exec(query, binaryID, jobID, reviewerID, revieweeID, rating, comment)
	if err != nil {
		return nil, fmt.Errorf("failed to create review: %w", err)
	}

	return &Review{ID: binaryID, JobID: jobID, ReviewerID: reviewerID, RevieweeID: revieweeID, Rating: rating, Comment: comment}, nil
}
