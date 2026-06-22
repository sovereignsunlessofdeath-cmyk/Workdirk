package services

import (
	"fmt"
	"workdirk/internal/repository"
)

type PaymentService struct {
	paymentRepo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{paymentRepo: repo}
}

func (s *PaymentService) HoldFunds(jobID, customerID, workerID []byte, amount float64) (*repository.Payment, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("escrow amount must be greater than zero")
	}

	p, err := s.paymentRepo.CreatePayment(jobID, customerID, workerID, amount)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PaymentService) ReleaseFunds(paymentID []byte) error {
	err := s.paymentRepo.UpdateStatus(paymentID, "Released")
	if err != nil {
		return fmt.Errorf("failed to release escrow: %w", err)
	}
	return nil
}
