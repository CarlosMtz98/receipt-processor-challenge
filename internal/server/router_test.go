package server

import (
	receiptHttp "github.com/CarlosMtz98/receipt-processor-challenge/internal/domain/receipt/delivery/http"
	"github.com/CarlosMtz98/receipt-processor-challenge/internal/domain/receipt/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckEndpoint(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReceiptService := mock.NewMockReceiptService(ctrl)
	receiptHandler := receiptHttp.NewReceiptHandler(mockReceiptService)
	// Create a new Gin router
	gin.SetMode(gin.TestMode)
	r := SetupRoutes(receiptHandler)

	// Create a test request to the /health endpoint
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
}
