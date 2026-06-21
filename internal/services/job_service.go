package services

import (
	"fmt"
	"workdirk/internal/repository"
)

type JobService struct {
	jobRepo *repository.JobRepository
}

func NewJobService(repo *repository.JobRepository) *JobService {
	return &JobService{jobRepo: repo}
}

func (s *JobService) CreateNewJob(customerID, workerID []byte, title, description string) (*repository.Job, error) {
	if title == "" || description == "" {
		return nil, fmt.Errorf("job title and description cannot be empty")
	}

	job, err := s.jobRepo.CreateJob(customerID, workerID, title, description)
	if err != nil {
		return nil, err
	}

	return job, nil
}
