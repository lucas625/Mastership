package service

import "net/http"

type Service interface {
	Add(http.ResponseWriter, *http.Request)
	Divide(http.ResponseWriter, *http.Request)
	Multiply(http.ResponseWriter, *http.Request)
	Subtract(http.ResponseWriter, *http.Request)
}
