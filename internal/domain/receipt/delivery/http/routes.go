package http

import "github.com/gin-gonic/gin"

func MapReceiptRoutes(routesGroup *gin.RouterGroup, handler ReceiptHandler) {
	routesGroup.POST("/process", handler.Create)
	routesGroup.GET("/:id/points", handler.GetPoints)
}
