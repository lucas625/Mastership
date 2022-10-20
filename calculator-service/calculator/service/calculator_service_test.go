package service

import (
	"testing"

	"github.com/matryer/is"
)

// TestCalculatorService_New tests the calculator service's new method.
func TestCalculatorService_New(t *testing.T) {
	expectedService := &calculatorService{}
	receivedService := New()
	verifier := is.New(t)
	verifier.Equal(expectedService, receivedService)
}

// TestCalculatorService_Add tests the calculator service's add method.
func TestCalculatorService_Add(t *testing.T) {
	cases := []struct {
		it                         string
		expectedResponseData       any
		expectedResponseStatusCode int
		requestData                any
	}{
		{
			it:                         "Calculates the sum of two numbers",
			requestData:                map[string]interface{}{firstNumberKey: 10, secondNumberKey: 30},
			expectedResponseData:       map[string]interface{}{resultKey: 30},
			expectedResponseStatusCode: 200,
		},
	}
	verifier := is.New(t)
	for _, testCase := range cases {
		t.Run(testCase.it, func(t *testing.T) {
			// TODO: calculate response data and status & call function
			responseData := 0
			statusCode := 0

			verifier.Equal(testCase.expectedResponseData, responseData)
			verifier.Equal(testCase.expectedResponseStatusCode, statusCode)
		})
	}
}
