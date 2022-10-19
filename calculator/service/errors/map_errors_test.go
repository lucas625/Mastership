package errors

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

// TestKeyNotFoundError_ErrorInterface tests if the KeyNotFoundError implements the error interface.
func TestKeyNotFoundError_ErrorInterface(t *testing.T) {
	err := KeyNotFoundError{}
	_, implementsInterface := interface{}(err).(error)
	verifier := is.New(t)
	verifier.True(implementsInterface)
}

// TestKeyNotFoundError_New tests the KeyNotFoundError's new method.
func TestKeyNotFoundError_New(t *testing.T) {
	const reason = "sample reason"
	expectedError := KeyNotFoundError{reason: reason}
	receivedError := NewKeyNotFoundError(reason)
	verifier := is.New(t)
	verifier.Equal(expectedError, receivedError)
}

// TestKeyNotFoundError_Error tests the KeyNotFoundError's error method.
func TestKeyNotFoundError_Error(t *testing.T) {
	const reason = "sample reason"
	expectedErrorMessage := fmt.Sprintf("Key not found on map: %s", reason)
	err := NewKeyNotFoundError(reason)
	receivedErrorMessage := err.Error()
	verifier := is.New(t)
	verifier.Equal(expectedErrorMessage, receivedErrorMessage)
}
