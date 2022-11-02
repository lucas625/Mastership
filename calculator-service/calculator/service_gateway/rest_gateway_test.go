package service_gateway

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/matryer/is"

	"github.com/lucas625/Mastership/calculator-service/calculator/mocks"
)

// TestRestGateway_New tests the restGateway's new method.
func TestRestGateway_New(t *testing.T) {
	serviceMock := &mocks.ServiceMock{}
	routerMock := &mocks.RouterMock{}
	serverMock := &mocks.ServerMock{}

	expectedGateway := restGateway{router: routerMock, server: serverMock, service: serviceMock}
	gateway := New(serviceMock, serverMock, routerMock)

	verifier := is.New(t)
	verifier.Equal(expectedGateway, gateway)
}

// TestRestGateway_CreateServer tests the restGateway's serve method.
func TestRestGateway_Serve(t *testing.T) {
	serviceMock := &mocks.ServiceMock{}
	routerMock := &mocks.RouterMock{
		HandleFuncFunc: func(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
			return nil
		},
	}
	serverMock := &mocks.ServerMock{
		ListenAndServeFunc: func() error {
			return nil
		},
	}

	gateway := New(serviceMock, serverMock, routerMock)
	err := gateway.Serve()

	verifier := is.New(t)
	verifier.Equal(nil, err)
}
