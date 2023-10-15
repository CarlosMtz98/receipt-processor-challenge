package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CarlosMtz98/receipt-processor-challenge/internal/domain/models"
	"github.com/CarlosMtz98/receipt-processor-challenge/internal/domain/receipt/mock"
	"github.com/CarlosMtz98/receipt-processor-challenge/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReceiptHandlerImpl_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReceipt := buildRandomReceipt(true, "Target")
	mockReceiptService := mock.NewMockReceiptService(ctrl)
	mockReceiptService.EXPECT().
		CreateReceipt(gomock.Any(), gomock.Any()).
		Return(&mockReceipt, nil)

	receiptHandler := NewReceiptHandler(mockReceiptService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/receipts/process", receiptHandler.Create)

	t.Run("Success", func(t *testing.T) {
		payload, err := json.Marshal(&mockReceipt)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.Code)

		var response = dto.CreateReceiptResponse{}
		if err := json.Unmarshal(resp.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}
		assert.NoError(t, err)
		assert.Equal(t, mockReceipt.ID.String(), response.ID)
	})

	t.Run("Malformed JSON", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, resp.Code)
		}
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("Invalid receipt schema", func(t *testing.T) {
		invalidReceipt := buildRandomReceipt(false, "")
		payload, err := json.Marshal(&invalidReceipt)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.NoError(t, err)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, resp.Code)
		}
		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestReceiptHandlerImpl_GetPoints(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReceipt := buildRandomReceipt(false, "Target")
	mockReceiptService := mock.NewMockReceiptService(ctrl)

	mockReceiptService.EXPECT().
		GetReceiptByID(gomock.Any(), mockReceipt.ID).
		Return(&mockReceipt, nil)

	mockReceiptService.EXPECT().
		GetReceiptPoints(gomock.Any(), &mockReceipt).
		Return(28, nil)

	receiptHandler := NewReceiptHandler(mockReceiptService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/receipts/:id/points", receiptHandler.GetPoints)

	t.Run("Success", func(t *testing.T) {
		url := fmt.Sprintf("/receipts/%s/points", mockReceipt.ID.String())
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.Code)
		}
		assert.Equal(t, http.StatusOK, resp.Code)

		var response = dto.GetPointsResponse{}
		if err := json.Unmarshal(resp.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 28, response.Points)
	})
}

func buildRandomReceipt(isNewReceipt bool, retailer string) models.Receipt {
	id := uuid.New()
	if isNewReceipt {
		id = uuid.Nil
	}
	receiptId := id
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
		Retailer:     retailer,
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items:        mockReceiptItems,
		Total:        "12.98",
	}
}
