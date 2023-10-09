package server

import (
	receiptHttp "github.com/CarlosMtz98/receipt-processor-challenge/internal/domain/receipt/delivery/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(receiptHandler receiptHttp.ReceiptHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := router.Group("/health")
	receipt := router.Group("/receipts")

	receiptHttp.MapReceiptRoutes(receipt, receiptHandler)

	health.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return router
}
