package calculator_requester

import "time"

//go:generate moq -out mocks/calculator_requester_mock.go -pkg mocks . CalculatorRequester

type CalculatorRequester interface {
	RequestAdd(data map[string]any) (map[string]any, error, time.Duration)
	RequestSubtract(data map[string]any) (map[string]any, error, time.Duration)
	RequestMultiply(data map[string]any) (map[string]any, error, time.Duration)
	RequestDivide(data map[string]any) (map[string]any, error, time.Duration)
}
