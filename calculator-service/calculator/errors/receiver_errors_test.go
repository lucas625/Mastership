package errors

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

// TestDecodeError_New tests the DecodeError's new method.
func TestDecodeError_New(t *testing.T) {
	expectedError := DecodeError{}
	receivedError := NewDecodeError()
	verifier := is.New(t)
	verifier.Equal(expectedError, receivedError)
}

// TestDecodeError_Error tests the DecodeError's error method.
func TestDecodeError_Error(t *testing.T) {
	expectedErrorMessage := fmt.Sprintf("Failed to decode")
	err := NewDecodeError()
	receivedErrorMessage := err.Error()
	verifier := is.New(t)
	verifier.Equal(expectedErrorMessage, receivedErrorMessage)
}

// TestUnmarshalError_New tests the UnmarshalError's new method.
func TestUnmarshalError_New(t *testing.T) {
	expectedError := UnmarshalError{}
	receivedError := NewUnmarshalError()
	verifier := is.New(t)
	verifier.Equal(expectedError, receivedError)
}

// TestUnmarshalError_Error tests the UnmarshalError's error method.
func TestUnmarshalError_Error(t *testing.T) {
	expectedErrorMessage := fmt.Sprintf("Failed to unmarshal")
	err := NewUnmarshalError()
	receivedErrorMessage := err.Error()
	verifier := is.New(t)
	verifier.Equal(expectedErrorMessage, receivedErrorMessage)
}
