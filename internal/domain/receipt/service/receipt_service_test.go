package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"receipt-processor-challenge/internal/domain/models"
	"receipt-processor-challenge/internal/domain/receipt/mock"
	"testing"
)

func TestReceiptServiceImpl_CreateReceipt(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReceiptRepo := mock.NewMockReceiptRepository(ctrl)
	service := NewReceiptService(mockReceiptRepo)

	// Create a test receipt
	receipt := &models.Receipt{
		ID: uuid.Nil, // ID is nil for this test
	}

	// Define the expected behavior of the mock repository
	mockReceiptRepo.EXPECT().Create(gomock.Any(), receipt).Times(1).Return(nil)

	// Call the CreateReceipt func
	createdReceipt, err := service.CreateReceipt(context.Background(), receipt)

	// Assert the results
	assert.Equal(t, receipt.ID, createdReceipt.ID)
	assert.NoError(t, err)
	assert.NotNil(t, createdReceipt)
	assert.NotEqual(t, uuid.Nil, createdReceipt.ID)
}

func TestReceiptServiceImpl_GetReceiptByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the repository
	mockReceiptRepo := mock.NewMockReceiptRepository(ctrl)
	receiptService := NewReceiptService(mockReceiptRepo)

	// Create a test receipt
	receipt := &models.Receipt{
		ID: uuid.New(),
	}

	// Define the expected behavior of the mock repository
	mockReceiptRepo.EXPECT().GetByID(gomock.Any(), receipt.ID).Times(1).Return(receipt, nil)

	// Call the GetReceiptByID func
	retrievedReceipt, err := receiptService.GetReceiptByID(context.Background(), receipt.ID)
	// Assert the results
	assert.Equal(t, receipt.ID, retrievedReceipt.ID)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedReceipt)
	assert.NotEqual(t, uuid.Nil, retrievedReceipt.ID)
}

func TestReceiptServiceImpl_GetReceiptPoints(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the repository
	mockReceiptRepo := mock.NewMockReceiptRepository(ctrl)
	receiptService := NewReceiptService(mockReceiptRepo)

	// Create a test receipt
	receipt := &models.Receipt{
		ID: uuid.New(),
	}

	// Define the expected behavior of the mock repository
	mockReceiptRepo.EXPECT().GetByID(gomock.Any(), receipt.ID).Times(1).Return(receipt, nil)

	// Call the GetReceiptPoints func
	points, err := receiptService.GetReceiptPoints(context.Background(), receipt.ID)
	// Assert the results
	assert.Equal(t, 100, points)
	assert.NoError(t, err)
	assert.NotNil(t, points)
}
