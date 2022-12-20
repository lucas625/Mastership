package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lucas625/Mastership/experimenter-service/experimenter/errors"
)

const (
	_allowedOperationsKey                    = "allowedOperations"
	_batchSizeKey                            = "batchSize"
	_interactionsKey                         = "interactions"
	_intervalBetweenBatchesInMillisecondsKey = "intervalBetweenBatchesInMilliseconds"
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
	batchSize, found := data[_batchSizeKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(_batchSizeKey)
	}
	interactions, found := data[_interactionsKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(_interactionsKey)
	}
	intervalBetweenBatchesInMilliseconds, found := data[_intervalBetweenBatchesInMillisecondsKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(_intervalBetweenBatchesInMillisecondsKey)
	}
	allowedOperations, found := data[_allowedOperationsKey]
	if !found {
		return nil, errors.NewKeyNotFoundError(_allowedOperationsKey)
	}

	batchSizeParsed, err := jsonAnyToInt(batchSize)
	if err != nil {
		return nil, err
	}
	interactionsParsed, err := jsonAnyToInt(interactions)
	if err != nil {
		return nil, err
	}
	intervalBetweenBatchesInMillisecondsParsed, err := jsonAnyToInt(intervalBetweenBatchesInMilliseconds)
	if err != nil {
		return nil, err
	}
	allowedOperationsParsed, err := jsonAnyToSliceOfStrings(allowedOperations)
	if err != nil {
		return nil, err
	}

	return newOperator(batchSizeParsed, interactionsParsed, intervalBetweenBatchesInMillisecondsParsed, allowedOperationsParsed)
}

func jsonAnyToInt(value any) (int, error) {
	valueFloat, ok := value.(float64)
	if !ok {
		return 0, errors.NewInvalidValueError(valueFloat, "int")
	}
	return int(valueFloat), nil
}

func jsonAnyToSliceOfStrings(value any) ([]string, error) {
	valueSlice, ok := value.([]string)
	if !ok {
		return nil, errors.NewInvalidValueError(valueSlice, "[]string")
	}
	return valueSlice, nil
}
