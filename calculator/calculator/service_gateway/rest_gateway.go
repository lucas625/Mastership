package service_gateway

import (
	"log"
	"net/http"

	"github.com/lucas625/Mastership/calculator/calculator"
)

type restGateway struct {
	router  calculator.Router
	server  calculator.Server
	service calculator.Service
}

var _ calculator.ServiceGateway = restGateway{}

func New(service calculator.Service, server calculator.Server, router calculator.Router) calculator.ServiceGateway {
	return restGateway{
		router:  router,
		server:  server,
		service: service,
	}
}

func (gateway restGateway) Serve() error {
	gateway.setEndpoints()
	log.Println("Server running!")
	return gateway.server.ListenAndServe()
}

func (gateway restGateway) setEndpoints() {
	log.Println("Setting endpoints")
	gateway.router.HandleFunc("/add", gateway.service.Add).Methods(http.MethodPost, http.MethodOptions)
	gateway.router.HandleFunc("/subtract", gateway.service.Subtract).Methods(http.MethodPost, http.MethodOptions)
	gateway.router.HandleFunc("/multiply", gateway.service.Multiply).Methods(http.MethodPost, http.MethodOptions)
	gateway.router.HandleFunc("/divide", gateway.service.Divide).Methods(http.MethodPost, http.MethodOptions)
}
