package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"receipt-processor-challenge/internal/domain/service"
	"receipt-processor-challenge/internal/dto"
)

type ReceiptHandler struct {
	ReceiptService service.ReceiptService
}

func CreateReceiptHandler(c *gin.Context) {
	var request dto.CreateReceiptRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := dto.ResponseErrorModel{
			Code:    1405,
			Message: "Could not parse the request body",
		}
		c.JSON(http.StatusBadRequest, response)
	}

	if err := validateRequestData(&request); err != nil {
		response := dto.ResponseErrorModel{
			Code:    1405,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := dto.CreateReceiptResponse{
		ID: uuid.New().String(),
	}

	c.JSON(http.StatusCreated, response)
}

func GetPointsHandler(c *gin.Context) {
	response := dto.GetPointsResponse{
		Points: 100,
	}

	c.JSON(http.StatusOK, response)
}
