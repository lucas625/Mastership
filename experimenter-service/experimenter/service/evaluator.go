package service

type evaluator struct {
	RTTSInMS []float64 `json:"rttsInMS"`
}

func newEvaluator(interactions int) *evaluator {
	return &evaluator{RTTSInMS: make([]float64, interactions)}
}
