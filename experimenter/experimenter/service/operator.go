package service

import (
	"github.com/lucas625/Mastership/experimenter-service/experimenter/errors"
)

type operator struct {
	AllowedOperations                    []string `json:"allowedOperations"`
	BatchSize                            int      `json:"batchSize"`
	Interactions                         int      `json:"interactions"`
	IntervalBetweenBatchesInMilliseconds int      `json:"intervalBetweenBatchesInMilliseconds"`
}

func newOperator(batchSize, interactions, intervalBetweenBatchesInMilliseconds int, allowedOperations []string) (*operator, error) {
	if batchSize <= 0 {
		return nil, errors.NewValidationError(batchSize, "batch size")
	}
	if interactions <= 0 {
		return nil, errors.NewValidationError(interactions, "interactions")
	}
	if intervalBetweenBatchesInMilliseconds < 0 {
		return nil, errors.NewValidationError(intervalBetweenBatchesInMilliseconds, "interval between batches in milliseconds")
	}
	if interactions < batchSize {
		return nil, errors.NewValidationError(batchSize, "batch size. Bigger than interactions")
	}
	allowedOperationsSize := len(allowedOperations)
	if allowedOperationsSize == 0 || allowedOperationsSize > 4 {
		return nil, errors.NewValidationError(allowedOperationsSize, "allowed operations size")
	}
	for _, operation := range allowedOperations {
		if operation != _sumOperation &&
			operation != _subtractOperation &&
			operation != _multiplyOperation &&
			operation != _divideOperation {
			return nil, errors.NewValidationError(operation, "invalid operation")
		}
	}
	return &operator{
			BatchSize:                            batchSize,
			Interactions:                         interactions,
			IntervalBetweenBatchesInMilliseconds: intervalBetweenBatchesInMilliseconds,
			AllowedOperations:                    allowedOperations,
		},
		nil
}
