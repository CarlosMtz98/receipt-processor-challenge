package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"receipt-processor-challenge/internal/domain/models"
	"sync"
)

var (
	// ErrReceiptNotFound is returned when a receipt is not found.
	ErrReceiptNotFound = errors.New("the receipt was not found in the repository")
	// ErrFailedToAddReceipt is returned when the receipt could not be added to the repository.
	ErrFailedToAddReceipt = errors.New("failed to add a new receipt to the repository")
)

type ReceiptRepository interface {
	Create(ctx context.Context, receipt *models.Receipt) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Receipt, error)
}

type InMemoryReceiptRepository struct {
	mu       sync.RWMutex
	receipts map[uuid.UUID]*models.Receipt
}

func InitReceiptRepository() ReceiptRepository {
	return &InMemoryReceiptRepository{
		receipts: make(map[uuid.UUID]*models.Receipt),
	}
}

func (memoryRepo *InMemoryReceiptRepository) Create(ctx context.Context, receipt *models.Receipt) error {
	if memoryRepo.receipts == nil {
		memoryRepo.mu.Lock()
		memoryRepo.receipts = make(map[uuid.UUID]*models.Receipt)
		memoryRepo.mu.Unlock()
	}

	if _, ok := memoryRepo.receipts[receipt.ID]; ok {
		return ErrFailedToAddReceipt
	}

	memoryRepo.mu.Lock()
	memoryRepo.receipts[receipt.ID] = receipt
	memoryRepo.mu.Unlock()
	return nil
}

func (memoryRepo *InMemoryReceiptRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Receipt, error) {
	memoryRepo.mu.RLock()
	defer memoryRepo.mu.RUnlock()

	if receipt, ok := memoryRepo.receipts[id]; ok {
		return receipt, nil
	}

	return nil, ErrReceiptNotFound
}
