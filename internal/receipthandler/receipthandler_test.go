package receipthandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestReceiptHandler_ProcessReceipt(t *testing.T) {
	// Create a test router and add the route
	router := gin.Default()
	handler := NewReceiptHandler()

	router.POST("/receipts/process", handler.ProcessReceipt)

	testCases := []struct {
		name             string
		requestBody      string
		expectedStatus   int
		expectedErrorMsg string
	}{
		{
			name:           "Valid Receipt",
			requestBody:    `{"retailer": "Test Store", "purchaseDate": "2022-01-01", "purchaseTime": "15:00", "items": [{"shortDescription": "Item A", "price": "5.00"}], "total": "5.00"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:             "Invalid Receipt Format",
			requestBody:      `{"invalid": "data"}`,
			expectedStatus:   http.StatusBadRequest,
			expectedErrorMsg: "The receipt is invalid",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "/receipts/process", strings.NewReader(tc.requestBody))
			resp := httptest.NewRecorder()

			// Serve the HTTP request
			router.ServeHTTP(resp, req)

			// Assert the status code
			assert.Equal(t, tc.expectedStatus, resp.Code, "Expected status code to match")

			if tc.expectedErrorMsg != "" {
				var responseBody map[string]string
				err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedErrorMsg, responseBody["error"])
			}
		})
	}
}

func TestReceiptHandler_GetPoints(t *testing.T) {
	// Create a test router and add the route
	router := gin.Default()
	handler := NewReceiptHandler()
	router.GET("/receipts/:id/points", handler.GetPoints)

	// First, store a receipt to test with
	router.POST("/receipts/process", handler.ProcessReceipt)
	testReceipt := `{"retailer": "Test Store", "purchaseDate": "2022-01-01", "purchaseTime": "15:00", "items": [{"shortDescription": "Item A", "price": "5.00"}], "total": "5.00"}`
	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", strings.NewReader(testReceipt))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var response map[string]string
	assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &response))
	receiptID, exists := response["id"]
	assert.True(t, exists, "Expected to receive a receipt ID")

	testCases := []struct {
		name             string
		receiptID        string
		expectedStatus   int
		expectedErrorMsg string
	}{
		{
			name:           "Valid Receipt ID",
			receiptID:      receiptID,
			expectedStatus: http.StatusOK,
		},
		{
			name:             "Invalid Receipt ID",
			receiptID:        "invalid_id",
			expectedStatus:   http.StatusNotFound,
			expectedErrorMsg: "No receipt found for that id",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/receipts/"+tc.receiptID+"/points", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			assert.Equal(t, tc.expectedStatus, resp.Code, "Expected status code to match")

			if tc.expectedErrorMsg != "" {
				var responseBody map[string]string
				err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedErrorMsg, responseBody["error"])
			}
		})
	}
}
