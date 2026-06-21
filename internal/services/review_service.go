package services

import (
	"fmt"
	"workdirk/internal/repository"
)

type ReviewService struct {
	reviewRepo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{reviewRepo: repo}
}

func (s *ReviewService) SubmitReview(jobID, reviewerID, revieweeID []byte, rating int, comment string) (*repository.Review, error) {
	if rating < 1 || rating > 5 {
		return nil, fmt.Errorf("rating score must be between 1 and 5")
	}

	r, err := s.reviewRepo.CreateReview(jobID, reviewerID, revieweeID, rating, comment)
	if err != nil {
		return nil, err
	}

	return r, nil
}
