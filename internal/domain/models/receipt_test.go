package models

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReceipt(t *testing.T) {
	receipt := Receipt{
		ID:           uuid.New(),
		Retailer:     "Sample Retailer",
		PurchaseDate: "2023-10-08",
		PurchaseTime: "13:01",
		Items: []ReceiptItem{
			{
				ShortDescription: "Item 1",
				Price:            "10.99",
			},
			{
				ShortDescription: "Item 2",
				Price:            "5.99",
			},
		},
		Total: "16.98",
	}

	t.Run("GetTotalAsFloat", func(t *testing.T) {
		expected := 16.98
		actual, err := receipt.GetTotalAsFloat()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("GetReceiptDateTime", func(t *testing.T) {
		expected := time.Date(2023, time.October, 8, 13, 1, 0, 0, time.UTC)
		actual, err := receipt.GetReceiptDatetime()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("GetReceiptItemsPriceAsFloat", func(t *testing.T) {
		receiptItem := receipt.Items[0]
		expected := 10.99
		assert.NotNilf(t, receiptItem, "ReceiptItem is nil")
		actual, err := receiptItem.GetPriceAsFloat()
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Valid receipt total", func(t *testing.T) {
		receipt := Receipt{
			ID:           uuid.New(),
			Retailer:     "Sample Retailer",
			PurchaseDate: "2023-10-08",
			PurchaseTime: "13:01",
			Items: []ReceiptItem{
				{
					ShortDescription: "Item 1",
					Price:            "10.99",
				},
				{
					ShortDescription: "Item 2",
					Price:            "5.99",
				},
			},
			Total: "16.98",
		}

		isValidReceipt, err := receipt.IsValid()
		assert.NoError(t, err)
		assert.Equal(t, true, isValidReceipt)
	})

	t.Run("Invalid recipt total, does not match with items price", func(t *testing.T) {
		receipt := Receipt{
			ID:           uuid.New(),
			Retailer:     "Sample Retailer",
			PurchaseDate: "2023-10-08",
			PurchaseTime: "13:01",
			Items: []ReceiptItem{
				{
					ShortDescription: "Item 1",
					Price:            "10.99",
				},
				{
					ShortDescription: "Item 2",
					Price:            "5.99",
				},
			},
			Total: "15.98",
		}
		isValidReceipt, err := receipt.IsValid()
		assert.NoError(t, err)
		assert.Equal(t, false, isValidReceipt)
	})
}
