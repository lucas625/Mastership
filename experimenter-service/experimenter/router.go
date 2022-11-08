package experimenter

import (
	"net/http"

	"github.com/gorilla/mux"
)

//go:generate moq -out mocks/router_mock.go -pkg mocks . Router

type Router interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
	ServeHTTP(http.ResponseWriter, *http.Request)
}

var _ Router = &mux.Router{}
