package errors

import "fmt"

type ValidationError struct {
	fieldName    string
	invalidValue any
}

var _ error = ValidationError{}

func NewValidationError(invalidValue any, fieldName string) error {
	return &ValidationError{fieldName: fieldName, invalidValue: invalidValue}
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("got an invalid value: %v for field: %s", err.invalidValue, err.fieldName)
}
