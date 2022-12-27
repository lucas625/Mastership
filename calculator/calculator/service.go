package calculator

import "net/http"

//go:generate moq -out mocks/service_mock.go -pkg mocks . Service

type Service interface {
	Add(http.ResponseWriter, *http.Request)
	Divide(http.ResponseWriter, *http.Request)
	Multiply(http.ResponseWriter, *http.Request)
	Subtract(http.ResponseWriter, *http.Request)
}
