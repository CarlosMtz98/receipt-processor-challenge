package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"receipt-processor-challenge/internal/domain/models"
	"receipt-processor-challenge/internal/domain/receipt/repository"
)

var (
	// ErrReceiptWithId is returned when a receipt is not found.
	ErrReceiptWithId = errors.New("the receipt was not found in the repository")
	// ErrMissingReceiptId is returned wh
	ErrMissingReceiptId = errors.New("the receipt id can't be null")
)

type ReceiptService interface {
	CreateReceipt(ctx context.Context, receipt *models.Receipt) (*models.Receipt, error)
	GetReceiptByID(ctx context.Context, receiptID uuid.UUID) (*models.Receipt, error)
	GetReceiptPoints(ctx context.Context, receiptID uuid.UUID) (int, error)
}

type ReceiptServiceImpl struct {
	receiptRepository repository.ReceiptRepository
}

func NewReceiptService(receiptRepository repository.ReceiptRepository) ReceiptService {
	return &ReceiptServiceImpl{
		receiptRepository: receiptRepository,
	}
}

func (s *ReceiptServiceImpl) CreateReceipt(ctx context.Context, receipt *models.Receipt) (*models.Receipt, error) {
	if receipt.ID != uuid.Nil {
		return nil, ErrReceiptWithId
	}

	receipt.ID = uuid.New()
	err := s.receiptRepository.Create(ctx, receipt)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (s *ReceiptServiceImpl) GetReceiptByID(ctx context.Context, id uuid.UUID) (*models.Receipt, error) {
	if id == uuid.Nil {
		return nil, ErrMissingReceiptId
	}

	receipt, err := s.receiptRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (s *ReceiptServiceImpl) GetReceiptPoints(ctx context.Context, id uuid.UUID) (int, error) {
	_, err := s.GetReceiptByID(ctx, id)
	if err != nil {
		return 0, err
	}

	return 100, nil
}
