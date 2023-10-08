package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"receipt-processor-challenge/internal/domain/models"
	"testing"
)

func TestInMemoryReceiptRepository_Create(t *testing.T) {
	// Initialize the receipt repository
	repo := InitReceiptRepository()

	// Create a test receipt the other fields are not required for this test
	receipt := &models.Receipt{
		ID: uuid.New(),
	}
	// Test create
	err := repo.Create(context.Background(), receipt)
	assert.NoError(t, err)

	// Test Create for an existing receipt
	err = repo.Create(context.Background(), receipt)
	assert.Error(t, err) // Expect an error
	assert.Contains(t, err.Error(), ErrFailedToAddReceipt.Error())
}

func TestInMemoryReceiptRepository_Create_ErrFailedToAddReceipt(t *testing.T) {
	// Initialize the repository
	repo := InitReceiptRepository()

	// Attempt to create a receipt with a duplicate ID
	receipt := &models.Receipt{
		ID: uuid.New(),
		// Initialize other fields as needed
	}

	// Create the receipt for the first time
	err := repo.Create(context.Background(), receipt)
	assert.NoError(t, err) // Expect no error

	// Attempt to create the same receipt again (should return ErrFailedToAddReceipt)
	err = repo.Create(context.Background(), receipt)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrFailedToAddReceipt.Error())
}

func TestInMemoryReceiptRepository_GetById(t *testing.T) {
	// Initialize the receipt repository
	repo := InitReceiptRepository()

	// Create a test receipt the other fields are not required for this test
	receipt := &models.Receipt{
		ID: uuid.New(),
	}

	// Test GetByID for an added receipt
	err := repo.Create(context.Background(), receipt)
	assert.NoError(t, err) // Expect no error

	retrievedReceipt, err := repo.GetByID(context.Background(), receipt.ID)
	assert.NoError(t, err) // Expect no error
	assert.NotNil(t, retrievedReceipt)
	assert.Equal(t, receipt.ID, retrievedReceipt.ID)
}

func TestInMemoryReceiptRepository_GetByID_ErrReceiptNotFound(t *testing.T) {
	// Initialize the repository
	repo := InitReceiptRepository()

	// Attempt to retrieve a non-existent receipt
	nonExistentID := uuid.New()
	retrievedReceipt, err := repo.GetByID(context.Background(), nonExistentID)

	// Expect an error and that the error message matches ErrReceiptNotFound
	assert.Error(t, err)
	assert.EqualError(t, err, ErrReceiptNotFound.Error())
	assert.Nil(t, retrievedReceipt)
}
