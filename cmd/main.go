package main

import (
	"log"
	"os"
	"receipt-processor-challenge/internal/domain/receipt/repository"
	"receipt-processor-challenge/internal/domain/receipt/service"
	"receipt-processor-challenge/internal/server/handler"
	"receipt-processor-challenge/internal/server/router"
)

func main() {
	receiptRepo := repository.InitReceiptRepository()
	receiptService := service.NewReceiptService(receiptRepo)
	receiptHandler := handler.NewReceiptHandler(receiptService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}

	r := router.SetupRoutes(receiptHandler)

	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
