package receiver

import (
	"encoding/json"
	"net/http"

	"github.com/lucas625/Mastership/calculator/service"
	"github.com/lucas625/Mastership/calculator/service/errors"
)

const (
	addOperation = iota
	subtractOperation
	multiplyOperation
	divideOperation
)

// An implementation of the service.Service interface.
type calculatorService struct{}

// New creates a new calculator service.
func New() service.Service {
	return &calculatorService{}
}

// Add gets a json with parameters described on the operator struct and adds them.
func (s calculatorService) Add(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, addOperation)
}

// Divide gets a json with parameters described on the operator struct and divides them.
func (s calculatorService) Divide(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, subtractOperation)
}

// Multiply gets a json with parameters described on the operator struct and multiplies them.
func (s calculatorService) Multiply(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, multiplyOperation)
}

// Subtract gets a json with parameters described on the operator struct and subtracts them.
func (s calculatorService) Subtract(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, divideOperation)
}

// processRequest processes the request.
func processRequest(responseWriter http.ResponseWriter, request *http.Request, operation int) {
	operator, err := parseRequest(responseWriter, request)
	if err == nil {
		result := processResult(operator, operation)
		sendResponse(responseWriter, result)
	}
}

// parseRequest parses the request.
func parseRequest(responseWriter http.ResponseWriter, request *http.Request) (*operator, error) {
	operator, err := requestToOperator(request)
	if err != nil {
		switch err.(type) {
		case errors.KeyNotFoundError:
			http.Error(responseWriter, err.Error(), 400)
		default:
			http.Error(responseWriter, err.Error(), 500)
		}
		return nil, err
	}
	return operator, nil
}

// processResult calculates the result based on the operation.
func processResult(operator *operator, operation int) float64 {
	var result float64
	switch operation {
	case addOperation:
		result = operator.firstNumber + operator.secondNumber
	case subtractOperation:
		result = operator.firstNumber - operator.secondNumber
	case multiplyOperation:
		result = operator.firstNumber * operator.secondNumber
	case divideOperation:
		result = operator.firstNumber / operator.secondNumber
	}
	return result
}

// sendResponse sends the response.
func sendResponse(responseWriter http.ResponseWriter, result float64) {
	response := response{result: result}
	responseAsBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
		return
	}
	_, err = responseWriter.Write(responseAsBytes)
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
	}
}
