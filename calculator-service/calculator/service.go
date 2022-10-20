package calculator

import "net/http"

//go:generate moq -out service_mock.go . Service

type Service interface {
	Add(http.ResponseWriter, *http.Request)
	Divide(http.ResponseWriter, *http.Request)
	Multiply(http.ResponseWriter, *http.Request)
	Subtract(http.ResponseWriter, *http.Request)
}
