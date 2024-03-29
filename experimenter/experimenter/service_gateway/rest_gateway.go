package service_gateway

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucas625/Mastership/experimenter-service/experimenter"
)

type restGateway struct {
	router  experimenter.Router
	server  experimenter.Server
	service experimenter.Service
}

var _ experimenter.ServiceGateway = restGateway{}

const _baseURL = "/api/experimenter"

func New(service experimenter.Service, server experimenter.Server, router experimenter.Router) experimenter.ServiceGateway {
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
	gateway.router.HandleFunc(fmt.Sprintf("%s/experiment", _baseURL), gateway.service.Experiment).Methods(http.MethodPost, http.MethodOptions)
}
