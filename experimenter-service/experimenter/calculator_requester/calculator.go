package calculator_requester

type CalculatorRequester struct {
	url string
}

func New(url string) *CalculatorRequester {
	return &CalculatorRequester{url: url}
}

// TODO: do request & add docs
func (cr *CalculatorRequester) RequestAdd(firstNumber, secondNumber float64) (float64, error) {
	// Do request
}

func (cr *CalculatorRequester) RequestSubtract(firstNumber, secondNumber float64) (float64, error) {
	// Do request
}

func (cr *CalculatorRequester) RequestMultiply(firstNumber, secondNumber float64) (float64, error) {
	// Do request
}

func (cr *CalculatorRequester) RequestDivide(firstNumber, secondNumber float64) (float64, error) {
	// Do request
}
