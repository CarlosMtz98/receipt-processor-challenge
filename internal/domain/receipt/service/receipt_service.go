package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"math"
	"receipt-processor-challenge/internal/domain/models"
	"receipt-processor-challenge/internal/domain/receipt/repository"
	"strings"
	"unicode"
)

var (
	// ErrReceiptWithId is returned when a receipt is not found.
	ErrReceiptWithId = errors.New("the receipt was not found in the repository")
	// ErrMissingReceiptId is returned when there is no id provided
	ErrMissingReceiptId = errors.New("the receipt id can't be null")
	// ErrReceiptIsNil  is returned when there is no receipt passed through the param
	ErrReceiptIsNil = errors.New("the receipt is null")
)

const (
	RoundAmountPoints      = 50
	TotalIsMultiplePoints  = 25
	ItemsPointsPerTwoItems = 5
	DateOddPoints          = 6
	TimePoints             = 10
)

type ReceiptService interface {
	CreateReceipt(ctx context.Context, receipt *models.Receipt) (*models.Receipt, error)
	GetReceiptByID(ctx context.Context, receiptID uuid.UUID) (*models.Receipt, error)
	GetReceiptPoints(ctx context.Context, receipt *models.Receipt) (int, error)
}

type ReceiptServiceImpl struct {
	receiptRepository repository.ReceiptRepository
}

func NewReceiptService(receiptRepository repository.ReceiptRepository) ReceiptService {
	return &ReceiptServiceImpl{
		receiptRepository: receiptRepository,
	}
}

func (s *ReceiptServiceImpl) CreateReceipt(ctx context.Context, receipt *models.Receipt) (*models.Receipt, error) {
	if receipt.ID != uuid.Nil {
		return nil, ErrReceiptWithId
	}

	receipt.ID = uuid.New()
	err := s.receiptRepository.Create(ctx, receipt)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (s *ReceiptServiceImpl) GetReceiptByID(ctx context.Context, id uuid.UUID) (*models.Receipt, error) {
	if id == uuid.Nil {
		return nil, ErrMissingReceiptId
	}

	receipt, err := s.receiptRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return receipt, nil
}

func (s *ReceiptServiceImpl) GetReceiptPoints(ctx context.Context, receipt *models.Receipt) (int, error) {
	if receipt == nil {
		return 0, ErrReceiptIsNil
	}
	if receipt == nil {
		return 0, ErrReceiptIsNil
	}

	points := 0
	points += countAlphanumerics(receipt.Retailer)
	points += isTotalRoundAmount(receipt)
	points += getReceiptTotalIfIsMultiplePoints(receipt)
	points += sumItemsPoints(receipt)
	points += getReceiptItemsDescriptionPoints(receipt)
	points += getDatePoints(receipt)
	points += getTimePoints(receipt)

	return points, nil
}

// 50 points if the total is a round dollar amount with no cents.
func isTotalRoundAmount(receipt *models.Receipt) int {
	receiptTotal, err := receipt.GetTotalAsFloat()
	if err != nil {
		return 0
	}
	if receiptTotal-math.Floor(receiptTotal+0.5) > 0 {
		return 0
	}

	return RoundAmountPoints
}

// One point for every alphanumeric character in the retailer name.
func countAlphanumerics(str string) int {
	count := 0

	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}

	return count
}

// 25 points if the total is a multiple of 0.25.
func getReceiptTotalIfIsMultiplePoints(receipt *models.Receipt) int {
	if receipt == nil {
		return 0
	}

	receiptTotal, err := receipt.GetTotalAsFloat()
	if err != nil {
		return 0
	}
	if isMultipleOf(receiptTotal, 0.25) {
		return TotalIsMultiplePoints
	}

	return 0
}

func isMultipleOf(a, b float64) bool {
	return math.Mod(a, b) == 0.0
}

// 5 points for every two items on the receipt.
func sumItemsPoints(receipt *models.Receipt) int {
	if receipt == nil || receipt.Items == nil {
		return 0
	}

	return len(receipt.Items) / 2 * ItemsPointsPerTwoItems
}

// 6 points if the day in the purchase date is odd.
func getDatePoints(receipt *models.Receipt) int {
	receiptDate, err := receipt.GetReceiptDatetime()
	if err != nil {
		return 0
	}

	if receiptDate.Day()%2 == 0 {
		return 0
	}

	return DateOddPoints
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func getTimePoints(receipt *models.Receipt) int {
	receiptDate, err := receipt.GetReceiptDatetime()
	if err != nil {
		return 0
	}

	receiptPurchaseHour := receiptDate.Hour() + 1

	if receiptPurchaseHour > 14 && receiptPurchaseHour < 16 {
		return TimePoints
	}

	return 0
}

// If the trimmed length of the item description is a multiple of 3,
// multiply the price by 0.2 and round up to the nearest integer.
// The result is the number of points earned.
func getReceiptItemsDescriptionPoints(receipt *models.Receipt) int {
	if receipt == nil {
		return 0
	}
	if receipt.Items == nil {
		return 0
	}

	points := 0

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, err := item.GetPriceAsFloat()
			if err != nil {
				// errMsg := fmt.Errorf("error parsing price for item: %v", err)
				return 0
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	return points
}
