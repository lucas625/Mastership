package service

import (
	"github.com/lucas625/Mastership/experimenter-service/experimenter/errors"
)

type operator struct {
	BatchSize                            int `json:"batchSize"`
	Interactions                         int `json:"interactions"`
	IntervalBetweenBatchesInMilliseconds int `json:"intervalBetweenBatchesInMilliseconds"`
}

func newOperator(batchSize, interactions, intervalBetweenBatchesInMilliseconds int) (*operator, error) {
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
	return &operator{BatchSize: batchSize, Interactions: interactions, IntervalBetweenBatchesInMilliseconds: intervalBetweenBatchesInMilliseconds}, nil
}
