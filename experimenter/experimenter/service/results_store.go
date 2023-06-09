package service

import "sync"

type resultsStore struct {
	mu                 sync.Mutex
	rttsInMicroseconds []float64
	failures           int
}

func newResultsStore(operator *operator) *resultsStore {
	return &resultsStore{rttsInMicroseconds: make([]float64, operator.Interactions)}
}

func (rs *resultsStore) addFailure() {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.failures++
}

func (rs *resultsStore) getFailures() int {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.failures
}

func (rs *resultsStore) setRTTAt(index int, rtt float64) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.rttsInMicroseconds[index] = rtt
}

func (rs *resultsStore) getRTTs() []float64 {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return rs.rttsInMicroseconds
}
