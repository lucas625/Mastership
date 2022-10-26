package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"

	"github.com/lucas625/Mastership/calculator-service/calculator/errors"
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
			expectedResponseData:       float64(40),
			expectedResponseStatusCode: 200,
			requestData:                map[string]any{firstNumberKey: 10, secondNumberKey: 30},
		},
		{
			it:                         "Unmarshal error",
			expectedResponseData:       errors.NewUnmarshalError().Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                "",
		},
		{
			it:                         "Invalid value error for first number",
			expectedResponseData:       errors.NewInvalidValueError("30", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: "30", secondNumberKey: 10},
		},
		{
			it:                         "Invalid value error for second number",
			expectedResponseData:       errors.NewInvalidValueError("10", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 30, secondNumberKey: "10"},
		},
		{
			it:                         "First number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(firstNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
		},
		{
			it:                         "Second number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(secondNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 10},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.it, func(t *testing.T) {
			verifier := is.New(t)
			dataAsBytes, err := json.Marshal(testCase.requestData)
			verifier.Equal(err, nil)

			request := httptest.NewRequest(http.MethodPost, "/add", bytes.NewReader(dataAsBytes))
			response := httptest.NewRecorder()

			calculatorService{}.Add(response, request)
			var data any
			err = json.Unmarshal(response.Body.Bytes(), &data)
			if err != nil {
				data = string(response.Body.Bytes())
			}

			verifier.Equal(testCase.expectedResponseData, data)
			verifier.Equal(testCase.expectedResponseStatusCode, response.Code)
		})
	}
}

// TestCalculatorService_Divide tests the calculator service's divide method.
func TestCalculatorService_Divide(t *testing.T) {
	cases := []struct {
		it                         string
		expectedResponseData       any
		expectedResponseStatusCode int
		requestData                any
	}{
		{
			it:                         "Calculates the division of two numbers",
			expectedResponseData:       float64(3),
			expectedResponseStatusCode: 200,
			requestData:                map[string]any{firstNumberKey: 30, secondNumberKey: 10},
		},
		{
			it:                         "Unmarshal error",
			expectedResponseData:       errors.NewUnmarshalError().Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                "",
		},
		{
			it:                         "Invalid value error for first number",
			expectedResponseData:       errors.NewInvalidValueError("30", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: "30", secondNumberKey: 10},
		},
		{
			it:                         "Invalid value error for second number",
			expectedResponseData:       errors.NewInvalidValueError("10", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 30, secondNumberKey: "10"},
		},
		{
			it:                         "First number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(firstNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
		},
		{
			it:                         "Second number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(secondNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 10},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.it, func(t *testing.T) {
			verifier := is.New(t)
			dataAsBytes, err := json.Marshal(testCase.requestData)
			verifier.Equal(err, nil)

			request := httptest.NewRequest(http.MethodPost, "/divide", bytes.NewReader(dataAsBytes))
			response := httptest.NewRecorder()

			calculatorService{}.Divide(response, request)
			var data any
			err = json.Unmarshal(response.Body.Bytes(), &data)
			if err != nil {
				data = string(response.Body.Bytes())
			}

			verifier.Equal(testCase.expectedResponseData, data)
			verifier.Equal(testCase.expectedResponseStatusCode, response.Code)
		})
	}
}

// TestCalculatorService_Multiply tests the calculator service's multiply method.
func TestCalculatorService_Multiply(t *testing.T) {
	cases := []struct {
		it                         string
		expectedResponseData       any
		expectedResponseStatusCode int
		requestData                any
	}{
		{
			it:                         "Calculates the multiplication of two numbers",
			expectedResponseData:       float64(300),
			expectedResponseStatusCode: 200,
			requestData:                map[string]any{firstNumberKey: 30, secondNumberKey: 10},
		},
		{
			it:                         "Unmarshal error",
			expectedResponseData:       errors.NewUnmarshalError().Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                "",
		},
		{
			it:                         "Invalid value error for first number",
			expectedResponseData:       errors.NewInvalidValueError("30", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: "30", secondNumberKey: 10},
		},
		{
			it:                         "Invalid value error for second number",
			expectedResponseData:       errors.NewInvalidValueError("10", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 30, secondNumberKey: "10"},
		},
		{
			it:                         "First number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(firstNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
		},
		{
			it:                         "Second number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(secondNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 10},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.it, func(t *testing.T) {
			verifier := is.New(t)
			dataAsBytes, err := json.Marshal(testCase.requestData)
			verifier.Equal(err, nil)

			request := httptest.NewRequest(http.MethodPost, "/multiply", bytes.NewReader(dataAsBytes))
			response := httptest.NewRecorder()

			calculatorService{}.Multiply(response, request)
			var data any
			err = json.Unmarshal(response.Body.Bytes(), &data)
			if err != nil {
				data = string(response.Body.Bytes())
			}

			verifier.Equal(testCase.expectedResponseData, data)
			verifier.Equal(testCase.expectedResponseStatusCode, response.Code)
		})
	}
}

// TestCalculatorService_Multiply tests the calculator service's multiply method.
func TestCalculatorService_Subtract(t *testing.T) {
	cases := []struct {
		it                         string
		expectedResponseData       any
		expectedResponseStatusCode int
		requestData                any
	}{
		{
			it:                         "Calculates the subtraction of two numbers",
			expectedResponseData:       float64(20),
			expectedResponseStatusCode: 200,
			requestData:                map[string]any{firstNumberKey: 30, secondNumberKey: 10},
		},
		{
			it:                         "Unmarshal error",
			expectedResponseData:       errors.NewUnmarshalError().Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                "",
		},
		{
			it:                         "Invalid value error for first number",
			expectedResponseData:       errors.NewInvalidValueError("30", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: "30", secondNumberKey: 10},
		},
		{
			it:                         "Invalid value error for second number",
			expectedResponseData:       errors.NewInvalidValueError("10", "float64").Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 30, secondNumberKey: "10"},
		},
		{
			it:                         "First number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(firstNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
		},
		{
			it:                         "Second number missing error",
			expectedResponseData:       errors.NewKeyNotFoundError(secondNumberKey).Error() + "\n",
			expectedResponseStatusCode: 400,
			requestData:                map[string]any{firstNumberKey: 10},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.it, func(t *testing.T) {
			verifier := is.New(t)
			dataAsBytes, err := json.Marshal(testCase.requestData)
			verifier.Equal(err, nil)

			request := httptest.NewRequest(http.MethodPost, "/subtract", bytes.NewReader(dataAsBytes))
			response := httptest.NewRecorder()

			calculatorService{}.Subtract(response, request)
			var data any
			err = json.Unmarshal(response.Body.Bytes(), &data)
			if err != nil {
				data = string(response.Body.Bytes())
			}

			verifier.Equal(testCase.expectedResponseData, data)
			verifier.Equal(testCase.expectedResponseStatusCode, response.Code)
		})
	}
}
