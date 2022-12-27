package calculator

import (
	"net/http"
)

//go:generate moq -out mocks/server_mock.go -pkg mocks . Server

type Server interface {
	ListenAndServe() error
}

var _ Server = &http.Server{}
