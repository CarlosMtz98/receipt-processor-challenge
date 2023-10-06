package repository

import (
	"context"
	"receipt-processor-challenge/internal/domain/model"
)

type ReceiptRepository interface {
	CreateReceipt(ctx context.Context, receipt *model.Receipt) (string, error)
	GetReceiptByID(ctx context.Context, id string) (*model.Receipt, error)
}
