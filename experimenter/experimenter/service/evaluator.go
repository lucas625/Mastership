package service

import (
	"github.com/montanaflynn/stats"
)

type evaluator struct {
	RTTSInMilliseconds []float64 `json:"rttsInMilliseconds"`
	Failures           int       `json:"failures"`
	Mean               float64   `json:"mean"`
	Median             float64   `json:"median"`
	StandardDeviation  float64   `json:"standardDeviation"`
	Workload           *operator `json:"workload"`
}

func newEvaluator(operator *operator) *evaluator {
	return &evaluator{Workload: operator}
}

func (e *evaluator) ProcessResults() error {
	data := stats.LoadRawData(e.RTTSInMilliseconds)
	var err error
	e.Mean, err = stats.Mean(data)
	if err != nil {
		return err
	}
	e.Median, err = stats.Median(data)
	if err != nil {
		return err
	}
	e.StandardDeviation, err = stats.StandardDeviation(data)
	return err
}
