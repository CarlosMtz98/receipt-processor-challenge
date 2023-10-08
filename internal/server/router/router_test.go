package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"receipt-processor-challenge/internal/domain/models"
	"receipt-processor-challenge/internal/domain/receipt/mock"
	"receipt-processor-challenge/internal/dto"
	"receipt-processor-challenge/internal/server/handler"
	"testing"
)

func TestHealthCheckEndpoint(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReceiptService := mock.NewMockReceiptService(ctrl)
	receiptHandler := handler.NewReceiptHandler(mockReceiptService)
	// Create a new Gin router
	r := SetupRoutes(receiptHandler)

	// Create a test request to the /health endpoint
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateReceiptEndpoint(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()
	// Create a new Gin router
	_, r := gin.CreateTestContext(w)

	mockReceiptService := mock.NewMockReceiptService(ctrl)
	receiptHandler := handler.NewReceiptHandler(mockReceiptService)

	// Create a new Gin router
	r.POST("/receipts/process", receiptHandler.Create)

	// Create a new HTTP request for the /receipts/process endpoint
	reqBody := []byte(`{"retailer": "Target", "purchaseDate": "2022-01-01", "purchaseTime": "13:01", "items": [{"shortDescription": "Mountain Dew 12PK", "price": "6.49"}], "total": "6.49"}`)
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	mockReceipt := buildRandomReceipt()

	mockReceiptService.EXPECT().
		CreateReceipt(gomock.Any(), gomock.Any()).
		Times(1).
		Return(&mockReceipt, nil)
	// Handle the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)

	var response = dto.CreateReceiptResponse{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	// Assert that the response ID matches the expected ID
	assert.Equal(t, mockReceipt.ID.String(), response.ID)
}

func TestGetPointsEndpoint(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()
	// Create a new Gin router
	_, r := gin.CreateTestContext(w)

	mockReceiptService := mock.NewMockReceiptService(ctrl)
	receiptHandler := handler.NewReceiptHandler(mockReceiptService)
	mockReceipt := buildRandomReceipt()

	r.GET("/receipts/:id/points", receiptHandler.GetPoints)

	// Create a new HTTP request for the /receipts/{id}/points endpoint
	url := fmt.Sprintf("/receipts/%s/points", mockReceipt.ID.String())
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	mockReceiptService.EXPECT().
		GetReceiptPoints(gomock.Any(), gomock.Any()).
		Times(1).
		Return(100, nil)
	// Handle the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	var response = dto.GetPointsResponse{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	// Assert that the response ID matches the expected ID
	assert.Equal(t, 100, response.Points)
}

func buildRandomReceipt() models.Receipt {
	receiptId := uuid.New()
	mockReceiptItems := []models.ReceiptItem{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            "6.49",
		},
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            "6.49",
		},
	}
	return models.Receipt{
		ID:           receiptId,
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items:        mockReceiptItems,
		Total:        "6.49",
	}
}
