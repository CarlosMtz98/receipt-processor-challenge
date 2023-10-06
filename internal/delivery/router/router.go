package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"receipt-processor-challenge/internal/delivery/handler"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Health check endpoint. receipt-processor-challenge API V1"})
	})

	r.POST("/receipts/process", handler.CreateReceiptHandler)
	r.GET("/receipts/:id/points", handler.GetPointsHandler)

	return r
}
