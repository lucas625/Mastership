package errors

import "fmt"

// KeyNotFoundError happens when trying to decode.
type KeyNotFoundError struct {
	reason string
}

// NewKeyNotFoundError creates a new KeyNotFoundError.
func NewKeyNotFoundError(reason string) error {
	return KeyNotFoundError{reason: reason}
}

// Error returns the string of the error.
func (err KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key not found on map: %s", err.reason)
}
