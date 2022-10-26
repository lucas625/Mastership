package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lucas625/Mastership/calculator-service/calculator/errors"
)

const (
	firstNumberKey  = "firstNumber"
	secondNumberKey = "secondNumber"
)

func requestToOperator(request *http.Request) (*operator, error) {
	bodyAsBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.NewDecodeError()
	}

	var data map[string]any
	err = json.Unmarshal(bodyAsBytes, &data)
	if err != nil {
		return nil, errors.NewUnmarshalError()
	}

	return mapToOperator(data)
}

func mapToOperator(data map[string]any) (*operator, error) {
	firstNumber, found := data[firstNumberKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(firstNumberKey)
	}
	secondNumber, found := data[secondNumberKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(secondNumberKey)
	}

	firstNumberParsed, ok := firstNumber.(float64)
	if !ok {
		return nil, errors.NewInvalidValueError(firstNumber, "float64")
	}
	secondNumberParsed, ok := secondNumber.(float64)
	if !ok {
		return nil, errors.NewInvalidValueError(secondNumber, "float64")
	}
	return &operator{firstNumber: firstNumberParsed, secondNumber: secondNumberParsed}, nil
}
