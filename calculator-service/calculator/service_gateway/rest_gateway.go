package service_gateway

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/lucas625/Mastership/calculator-service/calculator"
)

type restGateway struct {
	service       calculator.Service
	configuration *calculator.Configuration
}

var _ calculator.ServiceGateway = restGateway{}

func New(service calculator.Service, configuration *calculator.Configuration) calculator.ServiceGateway {
	return restGateway{service: service, configuration: configuration}
}

func (gateway restGateway) CreateServer() error {
	router := gateway.InitializeEndpoints()
	return gateway.runServer(router)
}

func (gateway restGateway) InitializeEndpoints() *mux.Router {
	log.Println("Setting endpoints")
	router := mux.NewRouter()
	router.HandleFunc("/add", gateway.service.Add)
	router.HandleFunc("/divide", gateway.service.Divide)
	router.HandleFunc("/multiply", gateway.service.Multiply)
	router.HandleFunc("/subtract", gateway.service.Subtract)
	return router
}

func (gateway restGateway) runServer(router *mux.Router) error {
	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", gateway.configuration.Port),
		WriteTimeout: 1 * time.Minute,
		ReadTimeout:  1 * time.Minute,
	}
	log.Println("Server running!")
	return server.ListenAndServe()
}
