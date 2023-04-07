// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"sync"
	
	"github.com/lucas625/Mastership/experimenter-service/experimenter/calculator_requester"
)

// Ensure, that CalculatorRequesterMock does implement calculator_requester.CalculatorRequester.
// If this is not the case, regenerate this file with moq.
var _ calculator_requester.CalculatorRequester = &CalculatorRequesterMock{}

// CalculatorRequesterMock is a mock implementation of calculator_requester.CalculatorRequester.
//
//	func TestSomethingThatUsesCalculatorRequester(t *testing.T) {
//
//		// make and configure a mocked calculator_requester.CalculatorRequester
//		mockedCalculatorRequester := &CalculatorRequesterMock{
//			RequestAddFunc: func(data map[string]any) (map[string]any, error, int64) {
//				panic("mock out the RequestAdd method")
//			},
//			RequestDivideFunc: func(data map[string]any) (map[string]any, error, int64) {
//				panic("mock out the RequestDivide method")
//			},
//			RequestMultiplyFunc: func(data map[string]any) (map[string]any, error, int64) {
//				panic("mock out the RequestMultiply method")
//			},
//			RequestSubtractFunc: func(data map[string]any) (map[string]any, error, int64) {
//				panic("mock out the RequestSubtract method")
//			},
//		}
//
//		// use mockedCalculatorRequester in code that requires calculator_requester.CalculatorRequester
//		// and then make assertions.
//
//	}
type CalculatorRequesterMock struct {
	// RequestAddFunc mocks the RequestAdd method.
	RequestAddFunc func(data map[string]any) (map[string]any, error, int64)

	// RequestDivideFunc mocks the RequestDivide method.
	RequestDivideFunc func(data map[string]any) (map[string]any, error, int64)

	// RequestMultiplyFunc mocks the RequestMultiply method.
	RequestMultiplyFunc func(data map[string]any) (map[string]any, error, int64)

	// RequestSubtractFunc mocks the RequestSubtract method.
	RequestSubtractFunc func(data map[string]any) (map[string]any, error, int64)

	// calls tracks calls to the methods.
	calls struct {
		// RequestAdd holds details about calls to the RequestAdd method.
		RequestAdd []struct {
			// Data is the data argument value.
			Data map[string]any
		}
		// RequestDivide holds details about calls to the RequestDivide method.
		RequestDivide []struct {
			// Data is the data argument value.
			Data map[string]any
		}
		// RequestMultiply holds details about calls to the RequestMultiply method.
		RequestMultiply []struct {
			// Data is the data argument value.
			Data map[string]any
		}
		// RequestSubtract holds details about calls to the RequestSubtract method.
		RequestSubtract []struct {
			// Data is the data argument value.
			Data map[string]any
		}
	}
	lockRequestAdd      sync.RWMutex
	lockRequestDivide   sync.RWMutex
	lockRequestMultiply sync.RWMutex
	lockRequestSubtract sync.RWMutex
}

// RequestAdd calls RequestAddFunc.
func (mock *CalculatorRequesterMock) RequestAdd(data map[string]any) (map[string]any, error, int64) {
	if mock.RequestAddFunc == nil {
		panic("CalculatorRequesterMock.RequestAddFunc: method is nil but CalculatorRequester.RequestAdd was just called")
	}
	callInfo := struct {
		Data map[string]any
	}{
		Data: data,
	}
	mock.lockRequestAdd.Lock()
	mock.calls.RequestAdd = append(mock.calls.RequestAdd, callInfo)
	mock.lockRequestAdd.Unlock()
	return mock.RequestAddFunc(data)
}

// RequestAddCalls gets all the calls that were made to RequestAdd.
// Check the length with:
//
//	len(mockedCalculatorRequester.RequestAddCalls())
func (mock *CalculatorRequesterMock) RequestAddCalls() []struct {
	Data map[string]any
} {
	var calls []struct {
		Data map[string]any
	}
	mock.lockRequestAdd.RLock()
	calls = mock.calls.RequestAdd
	mock.lockRequestAdd.RUnlock()
	return calls
}

// RequestDivide calls RequestDivideFunc.
func (mock *CalculatorRequesterMock) RequestDivide(data map[string]any) (map[string]any, error, int64) {
	if mock.RequestDivideFunc == nil {
		panic("CalculatorRequesterMock.RequestDivideFunc: method is nil but CalculatorRequester.RequestDivide was just called")
	}
	callInfo := struct {
		Data map[string]any
	}{
		Data: data,
	}
	mock.lockRequestDivide.Lock()
	mock.calls.RequestDivide = append(mock.calls.RequestDivide, callInfo)
	mock.lockRequestDivide.Unlock()
	return mock.RequestDivideFunc(data)
}

// RequestDivideCalls gets all the calls that were made to RequestDivide.
// Check the length with:
//
//	len(mockedCalculatorRequester.RequestDivideCalls())
func (mock *CalculatorRequesterMock) RequestDivideCalls() []struct {
	Data map[string]any
} {
	var calls []struct {
		Data map[string]any
	}
	mock.lockRequestDivide.RLock()
	calls = mock.calls.RequestDivide
	mock.lockRequestDivide.RUnlock()
	return calls
}

// RequestMultiply calls RequestMultiplyFunc.
func (mock *CalculatorRequesterMock) RequestMultiply(data map[string]any) (map[string]any, error, int64) {
	if mock.RequestMultiplyFunc == nil {
		panic("CalculatorRequesterMock.RequestMultiplyFunc: method is nil but CalculatorRequester.RequestMultiply was just called")
	}
	callInfo := struct {
		Data map[string]any
	}{
		Data: data,
	}
	mock.lockRequestMultiply.Lock()
	mock.calls.RequestMultiply = append(mock.calls.RequestMultiply, callInfo)
	mock.lockRequestMultiply.Unlock()
	return mock.RequestMultiplyFunc(data)
}

// RequestMultiplyCalls gets all the calls that were made to RequestMultiply.
// Check the length with:
//
//	len(mockedCalculatorRequester.RequestMultiplyCalls())
func (mock *CalculatorRequesterMock) RequestMultiplyCalls() []struct {
	Data map[string]any
} {
	var calls []struct {
		Data map[string]any
	}
	mock.lockRequestMultiply.RLock()
	calls = mock.calls.RequestMultiply
	mock.lockRequestMultiply.RUnlock()
	return calls
}

// RequestSubtract calls RequestSubtractFunc.
func (mock *CalculatorRequesterMock) RequestSubtract(data map[string]any) (map[string]any, error, int64) {
	if mock.RequestSubtractFunc == nil {
		panic("CalculatorRequesterMock.RequestSubtractFunc: method is nil but CalculatorRequester.RequestSubtract was just called")
	}
	callInfo := struct {
		Data map[string]any
	}{
		Data: data,
	}
	mock.lockRequestSubtract.Lock()
	mock.calls.RequestSubtract = append(mock.calls.RequestSubtract, callInfo)
	mock.lockRequestSubtract.Unlock()
	return mock.RequestSubtractFunc(data)
}

// RequestSubtractCalls gets all the calls that were made to RequestSubtract.
// Check the length with:
//
//	len(mockedCalculatorRequester.RequestSubtractCalls())
func (mock *CalculatorRequesterMock) RequestSubtractCalls() []struct {
	Data map[string]any
} {
	var calls []struct {
		Data map[string]any
	}
	mock.lockRequestSubtract.RLock()
	calls = mock.calls.RequestSubtract
	mock.lockRequestSubtract.RUnlock()
	return calls
}
