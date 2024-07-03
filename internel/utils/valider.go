package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidatePassport(fl validator.FieldLevel) bool {
	passportNumber := fl.Field().String()
	re := regexp.MustCompile(`^\d{4} \d{6}$`)
	return re.MatchString(passportNumber)
}
