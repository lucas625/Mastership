package calculator_requester

import (
	"fmt"

	"github.com/lucas625/Mastership/experimenter-service/experimenter/http_utils"
)

type requester struct {
	url string
}

func New(url string) CalculatorRequester {
	return &requester{url: url}
}

// RequestAdd requests the add method from the API.
func (cr *requester) RequestAdd(data map[string]any) (map[string]any, error) {
	return http_utils.PostRequest(fmt.Sprintf("%s/add", cr.url), data)
}

// RequestSubtract requests the subtract method from the API.
func (cr *requester) RequestSubtract(data map[string]any) (map[string]any, error) {
	return http_utils.PostRequest(fmt.Sprintf("%s/subtract", cr.url), data)
}

// RequestMultiply requests the multiply method from the API.
func (cr *requester) RequestMultiply(data map[string]any) (map[string]any, error) {
	return http_utils.PostRequest(fmt.Sprintf("%s/multiply", cr.url), data)
}

// RequestDivide requests the divide method from the API.
func (cr *requester) RequestDivide(data map[string]any) (map[string]any, error) {
	return http_utils.PostRequest(fmt.Sprintf("%s/divide", cr.url), data)
}
