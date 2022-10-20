package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lucas625/Mastership/calculator/calculator/errors"
)

const (
	firstNumberKey  = "firstNumber"
	resultKey       = "result"
	secondNumberKey = "secondNumber"
)

func requestToOperator(request *http.Request) (*operator, error) {
	bodyAsBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.NewDecodeError()
	}

	var data map[string]interface{}
	err = json.Unmarshal(bodyAsBytes, &data)
	if err != nil {
		return nil, errors.NewUnmarshalError()
	}

	return mapToOperator(data)
}

func mapToOperator(data map[string]interface{}) (*operator, error) {
	firstNumber, found := data[firstNumberKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(firstNumberKey)
	}
	secondNumber, found := data[secondNumberKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(secondNumberKey)
	}
	return &operator{firstNumber: firstNumber.(float64), secondNumber: secondNumber.(float64)}, nil
}
