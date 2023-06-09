package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lucas625/Mastership/calculator/calculator"
	"github.com/lucas625/Mastership/calculator/calculator/errors"
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
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	if request.Method == http.MethodOptions {
		return
	}
	operator, err := parseRequest(responseWriter, request)
	if err == nil {
		responseData, err := processResult(operator, operation)
		if err != nil {
			log.Println(err.Error())
			http.Error(responseWriter, err.Error(), 400)
			return
		}
		log.Println(fmt.Sprintf("%v %s %v = %v", operator.firstNumber, operation, operator.secondNumber, responseData.Result))
		sendResponse(responseWriter, responseData)
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
func processResult(operator *operator, operation string) (*responseData, error) {
	var result float64
	switch operation {
	case addOperation:
		result = operator.firstNumber + operator.secondNumber
	case subtractOperation:
		result = operator.firstNumber - operator.secondNumber
	case multiplyOperation:
		result = operator.firstNumber * operator.secondNumber
	case divideOperation:
		if operator.secondNumber == 0 {
			return nil, errors.NewZeroDivisionError()
		}
		result = operator.firstNumber / operator.secondNumber
	}
	return &responseData{Result: result}, nil
}

// sendResponse sends the response.
func sendResponse(responseWriter http.ResponseWriter, data *responseData) {
	responseAsBytes, err := json.Marshal(data)
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
