package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"receipt-processor-challenge/internal/domain/models"
	"receipt-processor-challenge/internal/domain/receipt/service"
	"receipt-processor-challenge/internal/dto"
	"receipt-processor-challenge/pkg/utils"
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
		utils.HandleBadRequest(c, "Could not parse the request body", err)
		return
	}

	if err := utils.ValidateStruct(c, receipt); err != nil {
		utils.HandleBadRequest(c, "The receipt params are not valid", err)
		return
	}

	createdReceipt, err := h.receiptSvc.CreateReceipt(c, receipt)
	if err != nil {
		utils.HandleInternalError(c, "Could not add new receipt", err)
		return
	}

	response := dto.CreateReceiptResponse{
		ID: createdReceipt.ID.String(),
	}

	c.JSON(http.StatusCreated, response)
	return
}

func (h ReceiptHandlerImpl) GetPoints(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		utils.HandleBadRequest(c, "Could not parse the request body", nil)
		return
	}

	receiptId, err := uuid.Parse(id)
	if err != nil {
		utils.HandleBadRequest(c, "Invalid ID format", err)
		return
	}

	receipt, err := h.receiptSvc.GetReceiptByID(c, receiptId)
	if err != nil {
		utils.HandleNotFound(c, fmt.Sprintf("Could not find the receipt with ID %s", receiptId))
		return
	}

	points, err := h.receiptSvc.GetReceiptPoints(c, receipt)
	if err != nil {
		utils.HandleInternalError(c, "Error calculating points", err)
		return
	}

	response := dto.GetPointsResponse{
		Points: points,
	}
	c.JSON(http.StatusOK, response)
	return
}
