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

	repo := mock.NewMockReceiptRepository(ctrl)
	service := NewReceiptService(repo)
	receipt := &models.Receipt{
		ID: uuid.Nil,
	}
	repo.EXPECT().
		Create(gomock.Any(), receipt).
		Return(nil)

	t.Run("Create Receipt", func(t *testing.T) {
		createdReceipt, err := service.CreateReceipt(context.Background(), receipt)
		assert.NoError(t, err)
		assert.NotNil(t, createdReceipt)
		assert.NotEqual(t, uuid.Nil, createdReceipt.ID)
		assert.Equal(t, receipt, createdReceipt)
	})
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
		ID:           uuid.New(),
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []models.ReceiptItem{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
	}

	// Call the GetReceiptPoints func
	points, err := receiptService.GetReceiptPoints(context.Background(), receipt)
	// Assert the results
	assert.Equal(t, 109, points)
	assert.NoError(t, err)
	assert.NotNil(t, points)
}

func TestIsTotalRoundAmount(t *testing.T) {
	// Create a test receipt with a round total amount (50 points)
	receipt := &models.Receipt{Total: "100.00"}
	points := isTotalRoundAmount(receipt)
	if points != 50 {
		t.Errorf("Expected 50 points, but got %d", points)
	}

	// Create a test receipt with a non-round total amount (0 points)
	receipt = &models.Receipt{Total: "99.99"}
	points = isTotalRoundAmount(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}

	receipt = &models.Receipt{Total: "1.50"}
	points = isTotalRoundAmount(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}

	receipt = &models.Receipt{Total: "0.50"}
	points = isTotalRoundAmount(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}
}

func TestCountAlphanumerics(t *testing.T) {
	// Test counting alphanumeric characters in a string
	str := "Hello123"
	count := countAlphanumerics(str)
	if count != len(str) {
		t.Errorf("Expected 7 alphanumeric characters, but got %d", count)
	}

	// Test counting non-alphanumeric characters in a string
	str = "!@#$%^&*"
	count = countAlphanumerics(str)
	if count != 0 {
		t.Errorf("Expected 0 alphanumeric characters, but got %d", count)
	}
}

func TestGetReceiptTotalIfIsMultiplePoints(t *testing.T) {
	// Create a test receipt with a total that is a multiple of 0.25 (25 points)
	receipt := &models.Receipt{Total: "25.00"}
	points := getReceiptTotalIfIsMultiplePoints(receipt)
	if points != 25 {
		t.Errorf("Expected 25 points, but got %d", points)
	}

	receipt = &models.Receipt{Total: "100.25"}
	points = getReceiptTotalIfIsMultiplePoints(receipt)
	if points != 25 {
		t.Errorf("Expected 25 points, but got %d", points)
	}

	// Create a test receipt with a total that is not a multiple of 0.25 (0 points)
	receipt = &models.Receipt{Total: "33.33"}
	points = getReceiptTotalIfIsMultiplePoints(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}
}

func TestGetReceiptDayPoints(t *testing.T) {
	// Create a test receipt with purchase day that is odd (6 points)
	receipt := &models.Receipt{
		PurchaseDate: "2023-10-03",
		PurchaseTime: "14:43",
	}
	points := getDatePoints(receipt)
	if points != 6 {
		t.Errorf("Expected 6 points, but got %d", points)
	}

	// Create a test receipt with purchase day that is even (0 points)
	receipt = &models.Receipt{
		PurchaseDate: "2023-10-08",
		PurchaseTime: "14:43",
	}
	points = getDatePoints(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}
}

func TestGetReceiptTimePoints(t *testing.T) {
	// Create a test receipt with purchase time that is after 2:00pm and before 4:00pm (10 points)
	receipt := &models.Receipt{
		PurchaseDate: "2023-10-03",
		PurchaseTime: "14:01",
	}
	points := getTimePoints(receipt)
	if points != 10 {
		t.Errorf("Expected 10 points, but got %d", points)
	}

	// Create a test receipt with purchase time that is before 4:00pm and before 4:00pm (10 points)
	receipt = &models.Receipt{
		PurchaseDate: "2023-10-03",
		PurchaseTime: "15:59",
	}
	points = getTimePoints(receipt)
	if points != 10 {
		t.Errorf("Expected 10 points, but got %d", points)
	}

	// Create a test receipt with purchase time that is before 2:00pm (0 points)
	receipt = &models.Receipt{
		PurchaseDate: "2023-10-08",
		PurchaseTime: "14:00",
	}
	points = getTimePoints(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}

	// Create a test receipt with purchase time that is after 4:00pm (0 points)
	receipt = &models.Receipt{
		PurchaseDate: "2023-10-08",
		PurchaseTime: "16:00",
	}
	points = getTimePoints(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}
}

func TestItemsDescriptionsPoints(t *testing.T) {
	receipt := &models.Receipt{
		Items: []models.ReceiptItem{
			{
				ShortDescription: "aaa",
				Price:            "10.00",
			},
			{
				ShortDescription: " aaa    ",
				Price:            "10.00",
			},
			{
				ShortDescription: " aaa b ccc  ",
				Price:            "10.00",
			},
			{
				ShortDescription: "aa",
				Price:            "10.00",
			},
			{
				ShortDescription: "  aa  ",
				Price:            "10.00",
			},
		},
	}
	points := getReceiptItemsDescriptionPoints(receipt)
	if points != 6 {
		t.Errorf("Expected 6 points, but got %d", points)
	}
}

func TestReceiptSumOfItemsPoints(t *testing.T) {
	// Create receipt to sum 5 points for every two items
	receipt := &models.Receipt{
		Items: []models.ReceiptItem{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
	}
	points := sumItemsPoints(receipt)
	if points != 10 {
		t.Errorf("Expected 10 points, but got %d", points)
	}

	// Create receipt to sum 5 points for every two items
	receipt = &models.Receipt{
		Items: []models.ReceiptItem{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
	}
	points = sumItemsPoints(receipt)
	if points != 0 {
		t.Errorf("Expected 0 points, but got %d", points)
	}
}
