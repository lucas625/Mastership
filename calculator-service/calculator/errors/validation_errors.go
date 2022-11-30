package errors

import (
	"fmt"
)

// ZeroDivisionError happens when trying to divide by 0
type ZeroDivisionError struct{}

var _ error = ZeroDivisionError{}

// NewZeroDivisionError creates a new ZeroDivisionError.
func NewZeroDivisionError() error {
	return ZeroDivisionError{}
}

// Error returns the string of the error.
func (err ZeroDivisionError) Error() string {
	return fmt.Sprintf("Request failed as it is not possible to divide by 0")
}
