package errors

import "fmt"

// KeyNotFoundError happens when trying to decode.
type KeyNotFoundError struct {
	text string
}

// NewKeyNotFoundError creates a new KeyNotFoundError.
func NewKeyNotFoundError(text string) error {
	return KeyNotFoundError{text: text}
}

// Error returns the string of the error.
func (err KeyNotFoundError) Error() string {
	return fmt.Sprintf("Key not found on map: %s", err.text)
}
