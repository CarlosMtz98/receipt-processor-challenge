package helpers

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// CurrencyValidator is a custom function to verify that a string matches the currency format
func CurrencyValidator(fl validator.FieldLevel) bool {
	currencyPattern := `^\d+\.\d{2}$`
	currencyRegex := regexp.MustCompile(currencyPattern)
	return currencyRegex.MatchString(fl.Field().String())
}
