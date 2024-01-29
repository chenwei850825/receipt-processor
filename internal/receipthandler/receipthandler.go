// Package receipthandler provides an HTTP handler interface for processing and retrieving points from receipts.
// It defines ReceiptHandler, which contains ReceiptProcessor for processing receipts.
// The package integrates with the Gin web framework for routing and handling HTTP requests.
package receipthandler

import (
	"net/http"
	"receipt-processor/pkg/receipt"

	"github.com/gin-gonic/gin"
)

// ReceiptHandler is the interface for handling receipt requests
type ReceiptHandler interface {
	ProcessReceipt(c *gin.Context)
	GetPoints(c *gin.Context)
}

type receiptHandler struct {
	receiptProcessor receipt.ReceiptProcessor
}

// NewReceiptHandler returns a new ReceiptHandler
func NewReceiptHandler() ReceiptHandler {
	receiptProcessor := receipt.NewReceiptProcessor()
	return &receiptHandler{
		receiptProcessor: receiptProcessor,
	}
}

// ProcessReceipt processes a receipt by calculating the points earned and storing the receipt
func (r receiptHandler) ProcessReceipt(c *gin.Context) {
	var receipt receipt.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
		return
	}

	id := r.receiptProcessor.StoreReceipt(receipt)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// GetPoints returns the points earned for a given receipt ID
func (r receiptHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")

	points, err := r.receiptProcessor.GetPoints(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
