package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationErrorItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(s interface{}) ([]ValidationErrorItem, error) {
	if err := validate.Struct(s); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			return formatValidationErrors(errs), nil
		}
		return nil, err
	}
	return nil, nil
}

func formatValidationErrors(errs validator.ValidationErrors) []ValidationErrorItem {
	var validationErrors []ValidationErrorItem
	for _, e := range errs {
		validationErrors = append(validationErrors, ValidationErrorItem{
			Field:   e.Field(),
			Message: messageForTag(e),
		})
	}
	return validationErrors
}

func messageForTag(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Minimum length is %s", e.Param())
	case "max":
		return fmt.Sprintf("Maximum length is %s", e.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", e.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", e.Param())
	case "gt":
		return fmt.Sprintf("Must be greater than %s", e.Param())
	case "lt":
		return fmt.Sprintf("Must be less than %s", e.Param())
	case "eqfield":
		return fmt.Sprintf("Must be equal to %s", e.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", e.Param())
	default:
		return "Invalid value"
	}
}
