package calculator_requester

//go:generate moq -out mocks/calculator_requester_mock.go -pkg mocks . CalculatorRequester
type CalculatorRequester interface {
	RequestAdd(data map[string]any) (map[string]any, error, int64)
	RequestSubtract(data map[string]any) (map[string]any, error, int64)
	RequestMultiply(data map[string]any) (map[string]any, error, int64)
	RequestDivide(data map[string]any) (map[string]any, error, int64)
}
