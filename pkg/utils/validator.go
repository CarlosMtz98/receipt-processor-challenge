package utils

import (
	"context"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	if err := validate.RegisterValidation("currency", CurrencyValidator); err != nil {
		return
	}
}

func ValidateStruct(ctx context.Context, s interface{}) error {
	return validate.StructCtx(ctx, s)
}
