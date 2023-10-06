package handler

import (
	"github.com/go-playground/validator/v10"
	"receipt-processor-challenge/internal/dto"
	"receipt-processor-challenge/internal/helpers"
)

// validateRequestData is a helper function to validate the request data using validation tags.
// Added to each member of the Request DTOs
func validateRequestData(request *dto.CreateReceiptRequest) error {
	validatorInstance := validator.New()
	if err := validatorInstance.RegisterValidation("currency", helpers.CurrencyValidator); err != nil {
		return err
	}

	if err := validatorInstance.Struct(request); err != nil {
		return err
	}

	return nil
}
