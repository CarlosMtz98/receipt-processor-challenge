package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type Receipt struct {
	ID           uuid.UUID
	Retailer     string        `json:"retailer" validate:"required"`
	PurchaseDate string        `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string        `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []ReceiptItem `json:"items" validate:"required,min=1,dive"`
	Total        string        `json:"total" validate:"required,currency"`
}

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,currency"`
}

func (r *Receipt) GetTotalAsFloat() (float64, error) {
	return strconv.ParseFloat(r.Total, 64)
}

func (r *Receipt) GetReceiptDatetime() (time.Time, error) {
	if r == nil {
		return time.Time{}, errors.New("receipt is nil")
	}

	if r.PurchaseDate == "" {
		return time.Time{}, errors.New("receipt date is nil or empty")
	}

	if r.PurchaseTime == "" {
		return time.Time{}, errors.New("receipt time is nil or empty")
	}

	receiptTimeStr := fmt.Sprintf("%s %s:00", r.PurchaseDate, r.PurchaseTime)
	date, err := time.Parse(time.DateTime, receiptTimeStr)

	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse receipt datetime: %v", err)
	}

	return date, nil
}

func (r *Receipt) IsValid() (bool, error) {
	totalPrice := 0.0
	for i := 0; i < len(r.Items); i++ {
		price, err := r.Items[i].GetReceiptItemPrice()
		if err != nil {
			return false, err
		}
		totalPrice += price
	}

	total, err := r.GetTotalAsFloat()
	if err != nil {
		return false, err
	}
	return total == totalPrice, nil
}

func (ri *ReceiptItem) GetPriceAsFloat() (float64, error) {
	return strconv.ParseFloat(ri.Price, 64)
}

func (r1 *ReceiptItem) GetReceiptItemPrice() (float64, error) {
	return strconv.ParseFloat(r1.Price, 64)
}
