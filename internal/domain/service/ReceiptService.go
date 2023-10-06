package service

import "receipt-processor-challenge/internal/domain/repository"

type ReceiptService struct {
	repo repository.ReceiptRepository
}
