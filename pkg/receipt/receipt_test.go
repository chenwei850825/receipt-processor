package receipt

import "testing"

func TestCalculatePoints(t *testing.T) {
	// Define a struct for test cases
	type testCase struct {
		name           string
		receipt        Receipt
		expectedPoints int
	}

	testCases := []testCase{
		{
			name: "Target Receipt",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
				},
				Total: "35.35",
			},
			expectedPoints: 28,
		},
		{
			name: "M&M Corner Market Receipt",
			receipt: Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Items: []Item{
					{"Gatorade", "2.25"},
					{"Gatorade", "2.25"},
					{"Gatorade", "2.25"},
					{"Gatorade", "2.25"},
				},
				Total: "9.00",
			},
			expectedPoints: 109,
		},
		{
			name: "Target Receipt: more retailer name characters",
			receipt: Receipt{
				Retailer:     "TargetTarget",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
				},
				Total: "35.35",
			},
			expectedPoints: 34,
		},
		{
			name: "Target Receipt: purchase date is even",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "13:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
				},
				Total: "35.35",
			},
			expectedPoints: 22,
		},
		{
			name: "Target Receipt: purchase time is between 2:00pm and 4:00pm",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "14:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
				},
				Total: "35.35",
			},
			expectedPoints: 38,
		},
		{
			name: "Target Receipt: total is a round dollar amount with no cents",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
				},
				Total: "35.00",
			},
			expectedPoints: 103,
		},
		{
			name: "Target Receipt: total is a multiple of 0.25",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
				},
				Total: "35.50",
			},
			expectedPoints: 53,
		},
		{
			name: "Target Receipt: 5 points for two items on the receipt",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
					{"Mountain Dew 12PK", "6.49"},
				},
				Total: "35.35",
			},
			expectedPoints: 33,
		},
		{
			name: "Target Receipt: trimmed item description is a multiple of 3 characters",
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []Item{
					{"Mountain Dew 12PK", "6.49"},
					{"Emils Cheese Pizza", "12.25"},
					{"Knorr Creamy Chicken", "1.26"},
					{"Doritos Nacho Cheese", "3.35"},
					{"   Klarbrunn 12-PK 12 FL OZ  ", "12.00"},
					{" additional test        ", "12.00"},
				},
				Total: "35.35",
			},
			expectedPoints: 36,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			points := calculatePoints(tc.receipt)
			if points != tc.expectedPoints {
				t.Errorf("Failed %s, expected %d points, got %d points", tc.name, tc.expectedPoints, points)
			}
		})
	}
}

func TestReceiptProcessor(t *testing.T) {
	// Define a struct for test cases
	type storeTestCase struct {
		receipt        Receipt
		expectedPoints int
	}

	processor := NewReceiptProcessor()

	storeTestCases := []storeTestCase{
		{
			receipt: Receipt{
				Retailer:     "Test Stores",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "15:00",
				Items: []Item{
					{"Item A", "5.00"},
					{"Item B", "10.00"},
				},
				Total: "15.00",
			},
			expectedPoints: 109,
		},
	}

	for _, tc := range storeTestCases {
		// Test StoreReceipt and GetPoints
		t.Run("StoreReceipt", func(t *testing.T) {
			id := processor.StoreReceipt(tc.receipt)
			if id == "" {
				t.Error("Failed to store receipt, received empty ID")
			}

			points, err := processor.GetPoints(id)
			if err != nil {
				t.Errorf("Error getting points for receipt: %v", err)
			}
			if points != tc.expectedPoints {
				t.Errorf("Expected %d points, got %d points", tc.expectedPoints, points)
			}
		})
	}

	// Testing GetPoints with invalid ID
	t.Run("GetPoints Invalid ID", func(t *testing.T) {
		_, err := processor.GetPoints("invalid_id")
		if err != ErrReceiptNotFound {
			t.Errorf("Expected ErrReceiptNotFound, got %v", err)
		}
	})
}
