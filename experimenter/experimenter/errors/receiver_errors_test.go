package errors

import (
	"fmt"
	"reflect"
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
	expectedErrorMessage := "Failed to decode"
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
	expectedErrorMessage := "Failed to unmarshal"
	err := NewUnmarshalError()
	receivedErrorMessage := err.Error()
	verifier := is.New(t)
	verifier.Equal(expectedErrorMessage, receivedErrorMessage)
}

// TestInvalidValueError_New tests the InvalidValueError's new method.
func TestInvalidValueError_New(t *testing.T) {
	const (
		invalidValue = "25"
		expectedType = "float64"
	)
	expectedError := InvalidValueError{invalidValue: invalidValue, expectedType: expectedType}
	receivedError := NewInvalidValueError(invalidValue, expectedType)
	verifier := is.New(t)
	verifier.Equal(expectedError, receivedError)
}

// TestInvalidValueError_Error tests the InvalidValueError's error method.
func TestInvalidValueError_Error(t *testing.T) {
	const (
		invalidValue = "25"
		expectedType = "float64"
	)

	expectedErrorMessage := fmt.Sprintf(
		"Value %v of type %v is not a %s",
		invalidValue,
		reflect.TypeOf(invalidValue),
		expectedType,
	)
	err := NewInvalidValueError(invalidValue, expectedType)
	receivedErrorMessage := err.Error()
	verifier := is.New(t)
	verifier.Equal(expectedErrorMessage, receivedErrorMessage)
}
