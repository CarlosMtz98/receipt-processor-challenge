package model

import "github.com/google/uuid"

type Receipt struct {
	ID               uuid.UUID
	Retailer         string
	PurchaseDateTime string
	Total            float64
	Items            []ReceiptItem
}

type ReceiptItem struct {
	ID               uuid.UUID
	ReceiptID        uuid.UUID
	ShortDescription string
	Price            float64
}
