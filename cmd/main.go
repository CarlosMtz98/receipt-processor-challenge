package main

import (
	"log"
	"os"
	receiptHttp "receipt-processor-challenge/internal/domain/receipt/delivery/http"
	"receipt-processor-challenge/internal/domain/receipt/repository"
	"receipt-processor-challenge/internal/domain/receipt/service"
	"receipt-processor-challenge/internal/server"
)

func main() {
	log.Println("Starting Server")
	receiptRepo := repository.InitReceiptRepository()
	receiptService := service.NewReceiptService(receiptRepo)
	receiptHandler := receiptHttp.NewReceiptHandler(receiptService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}

	r := server.SetupRoutes(receiptHandler)
	log.Printf("Server listening on port: %s", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
