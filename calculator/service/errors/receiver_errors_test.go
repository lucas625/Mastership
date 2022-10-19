package errors

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

// TestDecodeError_ErrorInterface tests if the DecodeError implements the error interface.
func TestDecodeError_ErrorInterface(t *testing.T) {
	err := DecodeError{}
	_, implementsInterface := interface{}(err).(error)
	verifier := is.New(t)
	verifier.True(implementsInterface)
}

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

// TestUnmarshalError_ErrorInterface tests if the UnmarshalError implements the error interface.
func TestUnmarshalError_ErrorInterface(t *testing.T) {
	err := UnmarshalError{}
	_, implementsInterface := interface{}(err).(error)
	verifier := is.New(t)
	verifier.True(implementsInterface)
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
