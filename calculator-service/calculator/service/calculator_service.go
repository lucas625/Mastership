package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lucas625/Mastership/calculator-service/calculator"
	"github.com/lucas625/Mastership/calculator-service/calculator/errors"
)

const (
	addOperation      = "+"
	subtractOperation = "-"
	multiplyOperation = "*"
	divideOperation   = "/"
)

// An implementation of the calculator.Service interface.
type calculatorService struct{}

var _ calculator.Service = calculatorService{}

// New creates a new calculator service.
func New() calculator.Service {
	return &calculatorService{}
}

// Add gets a json with parameters described on the operator struct and adds them.
func (s calculatorService) Add(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, addOperation)
}

// Divide gets a json with parameters described on the operator struct and divides them.
func (s calculatorService) Divide(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, divideOperation)
}

// Multiply gets a json with parameters described on the operator struct and multiplies them.
func (s calculatorService) Multiply(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, multiplyOperation)
}

// Subtract gets a json with parameters described on the operator struct and subtracts them.
func (s calculatorService) Subtract(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request, subtractOperation)
}

// processRequest processes the request.
func processRequest(responseWriter http.ResponseWriter, request *http.Request, operation string) {
	operator, err := parseRequest(responseWriter, request)
	if err == nil {
		result := processResult(operator, operation)
		log.Println(fmt.Sprintf("%v %s %v = %v", operator.firstNumber, operation, operator.secondNumber, result))
		sendResponse(responseWriter, result)
	}
}

// parseRequest parses the request.
func parseRequest(responseWriter http.ResponseWriter, request *http.Request) (*operator, error) {
	operator, err := requestToOperator(request)
	if err != nil {
		switch err.(type) {
		case errors.UnmarshalError:
			http.Error(responseWriter, err.Error(), 400)
		case errors.KeyNotFoundError:
			http.Error(responseWriter, err.Error(), 400)
		case errors.DecodeError:
			http.Error(responseWriter, err.Error(), 400)
		case errors.InvalidValueError:
			http.Error(responseWriter, err.Error(), 400)
		default:
			http.Error(responseWriter, err.Error(), 500)
		}
		log.Println(err)
		return nil, err
	}
	return operator, nil
}

// processResult calculates the result based on the operation.
func processResult(operator *operator, operation string) float64 {
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
	responseAsBytes, err := json.Marshal(result)
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
		log.Println(err.Error())
		return
	}
	_, err = responseWriter.Write(responseAsBytes)
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
		log.Println(err.Error())
		return
	}
}
