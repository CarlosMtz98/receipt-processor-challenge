package utils

import (
	"github.com/CarlosMtz98/receipt-processor-challenge/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleBadRequest(c *gin.Context, message string, err error) {
	c.JSON(http.StatusBadRequest, dto.ResponseErrorModel{
		Code:    http.StatusBadRequest,
		Message: message,
		Details: err.Error(),
	})
}

func HandleNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, dto.ResponseErrorModel{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

func HandleInternalError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, dto.ResponseErrorModel{
		Code:    http.StatusInternalServerError,
		Message: message,
		Details: err.Error(),
	})
}
