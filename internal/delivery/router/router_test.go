package router

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckEndpoint(t *testing.T) {
	// Create a new Gin router
	r := SetupRoutes()

	// Create a test request to the /health endpoint
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateReceiptEndpoint(t *testing.T) {
	// Create a new Gin router
	r := SetupRoutes()

	// Create a new HTTP request for the /receipts/process endpoint
	reqBody := []byte(`{"retailer": "Target", "purchaseDate": "2022-01-01", "purchaseTime": "13:01", "items": [{"shortDescription": "Mountain Dew 12PK", "price": "6.49"}], "total": "6.49"}`)
	req, err := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Handle the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetPointsEndpoint(t *testing.T) {
	// Create a new Gin router
	r := SetupRoutes()

	// Create a new HTTP request for the /receipts/{id}/points endpoint
	req, err := http.NewRequest("GET", "/receipts/some-id/points", nil)
	assert.NoError(t, err)

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()

	// Handle the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
}
