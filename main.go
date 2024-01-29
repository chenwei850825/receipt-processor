package main

import (
	"receipt-processor/internal/receipthandler"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	rh := receipthandler.NewReceiptHandler()
	router.POST("/receipts/process", rh.ProcessReceipt)
	router.GET("/receipts/:id/points", rh.GetPoints)

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
