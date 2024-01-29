package receipt

// Item represents an item on a receipt.
type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required,numeric"`
}

// Receipt represents a receipt.
type Receipt struct {
	Retailer     string `json:"retailer" binding:"required"` // Non-whitespace characters
	PurchaseDate string `json:"purchaseDate" binding:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" binding:"required,datetime=15:04"`
	Items        []Item `json:"items" binding:"required,dive,required"` // Dive: apply validation to each element in the slice
	Total        string `json:"total" binding:"required"`               // Numeric with two decimal places
}

// ProcessedReceipt represents a processed receipt with calculated points.
type ProcessedReceipt struct {
	Receipt Receipt
	Points  int
}
