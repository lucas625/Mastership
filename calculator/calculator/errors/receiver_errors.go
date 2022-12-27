package errors

import (
	"fmt"
	"reflect"
)

// DecodeError happens when trying to decode.
type DecodeError struct{}

var _ error = DecodeError{}

// NewDecodeError creates a new DecodeError.
func NewDecodeError() error {
	return DecodeError{}
}

// Error returns the string of the error.
func (err DecodeError) Error() string {
	return "Failed to decode"
}

// UnmarshalError happens when trying to unmarshal.
type UnmarshalError struct{}

var _ error = UnmarshalError{}

// NewUnmarshalError creates a new UnmarshalError.
func NewUnmarshalError() error {
	return UnmarshalError{}
}

// Error returns the string of the error.
func (err UnmarshalError) Error() string {
	return "Failed to unmarshal"
}

// InvalidValueError happens when trying to parse an invalid value.
type InvalidValueError struct {
	expectedType string
	invalidValue any
}

var _ error = InvalidValueError{}

// NewInvalidValueError creates a new InvalidValueError.
func NewInvalidValueError(invalidValue any, expectedType string) error {
	return InvalidValueError{invalidValue: invalidValue, expectedType: expectedType}
}

// Error returns the string of the error.
func (err InvalidValueError) Error() string {
	return fmt.Sprintf(
		"Value %v of type %v is not a %s",
		err.invalidValue,
		reflect.TypeOf(err.invalidValue),
		err.expectedType,
	)
}
