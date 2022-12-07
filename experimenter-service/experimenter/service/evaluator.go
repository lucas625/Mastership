package service

import "github.com/montanaflynn/stats"

type evaluator struct {
	RTTSInMicroseconds []int64   `json:"rttsInMicroseconds"`
	Failures           int       `json:"failures"`
	Mean               float64   `json:"mean"`
	StandardDeviation  float64   `json:"standardDeviation"`
	Workload           *operator `json:"workload"`
}

func newEvaluator(operator *operator) *evaluator {
	return &evaluator{RTTSInMicroseconds: make([]int64, operator.Interactions), Workload: operator}
}

func (e *evaluator) ProcessResults() error {
	data := stats.LoadRawData(e.RTTSInMicroseconds)
	var err error
	e.Mean, err = stats.Mean(data)
	if err != nil {
		return err
	}
	e.StandardDeviation, err = stats.StandardDeviation(data)
	return err
}
