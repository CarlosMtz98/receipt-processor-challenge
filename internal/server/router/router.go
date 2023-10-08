package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"receipt-processor-challenge/internal/server/handler"
)

func SetupRoutes(receiptHandler handler.ReceiptHandler) *gin.Engine {
	r := gin.Default()

	health := r.Group("/health")
	receipt := r.Group("/receipts")

	health.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	receipt.POST("/process", receiptHandler.Create)
	receipt.GET("/:id/points", receiptHandler.GetPoints)

	return r
}
