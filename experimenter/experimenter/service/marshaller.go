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

	err = validateMapToOperator(data)
	if err != nil {
		return nil, err
	}

	var opt operator
	err = json.Unmarshal(bodyAsBytes, &opt)
	if err != nil {
		return nil, errors.NewUnmarshalError()
	}

	return newOperator(opt.BatchSize, opt.Interactions, opt.IntervalBetweenBatchesInMilliseconds, opt.AllowedOperations)
}

func validateMapToOperator(data map[string]any) error {
	_, found := data[_batchSizeKey]
	if !found {
		return errors.NewKeyNotFoundError(_batchSizeKey)
	}
	_, found = data[_interactionsKey]
	if !found {
		return errors.NewKeyNotFoundError(_interactionsKey)
	}
	_, found = data[_intervalBetweenBatchesInMillisecondsKey]
	if !found {
		return errors.NewKeyNotFoundError(_intervalBetweenBatchesInMillisecondsKey)
	}
	_, found = data[_allowedOperationsKey]
	if !found {
		return errors.NewKeyNotFoundError(_allowedOperationsKey)
	}
	return nil
}
