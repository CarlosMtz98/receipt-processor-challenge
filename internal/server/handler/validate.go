package handler

import (
	"github.com/go-playground/validator/v10"
	"receipt-processor-challenge/internal/domain/models"
	"receipt-processor-challenge/internal/helpers"
)

// validateRequestData is a helper function to validate the request data using validation tags.
// Added to each member of the Request DTOs
func validateRequestData(receipt models.Receipt) error {
	validatorInstance := validator.New()
	if err := validatorInstance.RegisterValidation("currency", helpers.CurrencyValidator); err != nil {
		return err
	}

	if err := validatorInstance.Struct(receipt); err != nil {
		return err
	}

	return nil
}
