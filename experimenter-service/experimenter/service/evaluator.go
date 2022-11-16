package service

type evaluator struct {
	RTTSInMS []int64 `json:"rttsInMS"`
	Failures int     `json:"failures"`
}

func newEvaluator(interactions int) *evaluator {
	return &evaluator{RTTSInMS: make([]int64, interactions)}
}
