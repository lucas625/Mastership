package main

import (
	"fmt"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/calculator_requester"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/lucas625/Mastership/experimenter-service/experimenter"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/configuration"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/service"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/service_gateway"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config, err := configuration.New()
	if err != nil {
		return err
	}
	gateway := setupGateway(config)
	return gateway.Serve()
}

func setupGateway(config *configuration.Configuration) experimenter.ServiceGateway {
	requester := calculator_requester.New(config.CalculatorServiceAddress)
	experimenterService := service.New(requester)
	router := mux.NewRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", config.Port),
		WriteTimeout: 1 * time.Minute,
		ReadTimeout:  1 * time.Minute,
	}
	return service_gateway.New(experimenterService, server, router)
}
