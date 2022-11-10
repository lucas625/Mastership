package service

type operator struct {
	batchSize                            int
	interactions                         int
	intervalBetweenBatchesInMilliseconds int
}

func newOperator(batchSize, interactions, intervalBetweenBatchesInMilliseconds int) *operator {
	return &operator{batchSize: batchSize, interactions: interactions, intervalBetweenBatchesInMilliseconds: intervalBetweenBatchesInMilliseconds}
}
