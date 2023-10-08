package main

import (
	"receipt-processor-challenge/internal/domain/receipt/repository"
	"receipt-processor-challenge/internal/domain/receipt/service"
	"receipt-processor-challenge/internal/server/handler"
	"receipt-processor-challenge/internal/server/router"
)

func main() {
	receiptRepo := repository.InitReceiptRepository()
	receiptService := service.NewReceiptService(receiptRepo)
	receiptHandler := handler.NewReceiptHandler(receiptService)

	r := router.SetupRoutes(receiptHandler)

	r.Run(":8080")
}
