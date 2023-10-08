package models

import "github.com/google/uuid"

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
