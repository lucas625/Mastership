package service

type operator struct {
	firstNumber  float64
	secondNumber float64
}

func newOperator(firstNumber, secondNumber float64) *operator {
	return &operator{firstNumber: firstNumber, secondNumber: secondNumber}
}
