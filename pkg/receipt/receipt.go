// Package receipt provides a ReceiptProcessor interface to handle receipt processing logic.
// It includes functionality for calculating points based on specific rules,
// storing processed receipts, and retrieving calculated points for given receipts ids.
package receipt

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	receiptStore = make(map[string]ProcessedReceipt)
	storeMutex   sync.RWMutex
	// ErrReceiptNotFound is returned when a receipt ID is not found
	ErrReceiptNotFound = errors.New("Receipt not found")
)

// ReceiptProcessor is the interface for calculating points, storing receipts, and retrieving points
type ReceiptProcessor interface {
	StoreReceipt(receipt Receipt) string
	GetPoints(id string) (int, error)
}

type receiptProcessor struct{}

// NewReceiptProcessor returns a new ReceiptProcessor
func NewReceiptProcessor() ReceiptProcessor {
	return &receiptProcessor{}
}

// StoreReceipt processes a receipt by calculating the points earned and storing the receipt
// returns the generated ID of the receipt
func (r receiptProcessor) StoreReceipt(receipt Receipt) string {
	points := calculatePoints(receipt)
	id := generateReceiptID(receipt)
	storeMutex.Lock()
	receiptStore[id] = ProcessedReceipt{Receipt: receipt, Points: points}
	storeMutex.Unlock()
	return id
}

// GetPoints returns the points earned for a given receipt ID
func (r receiptProcessor) GetPoints(id string) (int, error) {
	storeMutex.RLock()
	processed, exists := receiptStore[id]
	storeMutex.RUnlock()
	if !exists {
		return 0, ErrReceiptNotFound
	}
	return processed.Points, nil
}

func generateReceiptID(receipt Receipt) string {
	return uuid.New().String()
}

func calculatePoints(receipt Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	reg := regexp.MustCompile("[0-9A-Za-z]")
	points += len(reg.FindAllString(receipt.Retailer, -1))

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && int(total*100)%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for item descriptions and prices.
	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && date.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil {
		if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
			points += 10
		}
	}

	return points
}
