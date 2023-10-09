package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"receipt-processor-challenge/internal/domain/models"
	"sync"
	"testing"
)

func TestInMemoryReceiptRepository(t *testing.T) {
	t.Run("Create and GetByID", func(t *testing.T) {
		repo := InitReceiptRepository().(*InMemoryReceiptRepository)

		receipt := &models.Receipt{
			ID: uuid.New(),
		}

		err := repo.Create(context.Background(), receipt)
		assert.NoError(t, err)

		retrievedReceipt, err := repo.GetByID(context.Background(), receipt.ID)
		assert.NoError(t, err)
		assert.Equal(t, receipt, retrievedReceipt)
	})

	t.Run("Create Duplicate", func(t *testing.T) {
		repo := InitReceiptRepository().(*InMemoryReceiptRepository)

		receipt := &models.Receipt{
			ID: uuid.New(),
		}

		err := repo.Create(context.Background(), receipt)
		assert.NoError(t, err)

		err = repo.Create(context.Background(), receipt)
		assert.Error(t, err)
		assert.Equal(t, ErrFailedToAddReceipt, err)
	})
}

func TestInMemoryReceiptRepository_EdgeCases(t *testing.T) {
	t.Run("Create with In Memory Nil Map", func(t *testing.T) {
		repo := &InMemoryReceiptRepository{
			mu:       sync.RWMutex{},
			receipts: nil,
		}

		receipt := &models.Receipt{
			ID: uuid.New(),
		}

		err := repo.Create(context.Background(), receipt)
		assert.NoError(t, err)

		retrievedReceipt, err := repo.GetByID(context.Background(), receipt.ID)
		assert.NoError(t, err)
		assert.Equal(t, receipt, retrievedReceipt)
	})

	t.Run("GetByID from Nil Map", func(t *testing.T) {
		repo := &InMemoryReceiptRepository{
			mu:       sync.RWMutex{},
			receipts: nil,
		}

		nonExistentID := uuid.New()
		retrievedReceipt, err := repo.GetByID(context.Background(), nonExistentID)
		assert.Error(t, err)
		assert.Nil(t, retrievedReceipt)
		assert.Equal(t, ErrReceiptNotFound, err)
	})
}
