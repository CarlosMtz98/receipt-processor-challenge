package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"receipt-processor-challenge/internal/domain/models"
	"receipt-processor-challenge/internal/domain/receipt/service"
	"receipt-processor-challenge/internal/dto"
)

type ReceiptHandler interface {
	Create(c *gin.Context)
	GetPoints(c *gin.Context)
}

type ReceiptHandlerImpl struct {
	receiptSvc service.ReceiptService
}

func NewReceiptHandler(receiptService service.ReceiptService) ReceiptHandler {
	return &ReceiptHandlerImpl{
		receiptSvc: receiptService,
	}
}

func (h ReceiptHandlerImpl) Create(c *gin.Context) {
	receipt := &models.Receipt{}

	if err := c.Bind(&receipt); err != nil {
		response := dto.ResponseErrorModel{
			Code:    1405,
			Message: "Could not parse the request body",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	createdReceipt, err := h.receiptSvc.CreateReceipt(c, receipt)
	if err != nil {
		response := dto.ResponseErrorModel{
			Code:    1405,
			Message: "Could not add new receipt",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := dto.CreateReceiptResponse{
		ID: createdReceipt.ID.String(),
	}

	c.JSON(http.StatusCreated, response)
}

func (h ReceiptHandlerImpl) GetPoints(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		response := dto.ResponseErrorModel{
			Code:    1405,
			Message: "Could not parse the request body",
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Convert 'idStr' to a UUID
	receiptId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	points, err := h.receiptSvc.GetReceiptPoints(c, receiptId)
	if err != nil {
		errMsg := fmt.Sprintf("Could not found receipt with id: %s", id)
		response := dto.ResponseErrorModel{
			Code:    1404,
			Message: errMsg,
		}
		c.JSON(http.StatusNotFound, response)
		return
	}
	response := dto.GetPointsResponse{
		Points: points,
	}

	c.JSON(http.StatusOK, response)
	return
}
