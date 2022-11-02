package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lucas625/Mastership/calculator-service/calculator"
	"log"
	"net/http"
	"time"

	"github.com/lucas625/Mastership/calculator-service/calculator/configuration"
	"github.com/lucas625/Mastership/calculator-service/calculator/service"
	"github.com/lucas625/Mastership/calculator-service/calculator/service_gateway"
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

func setupGateway(config *configuration.Configuration) calculator.ServiceGateway {
	calculatorService := service.New()
	router := mux.NewRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", config.Port),
		WriteTimeout: 1 * time.Minute,
		ReadTimeout:  1 * time.Minute,
	}
	return service_gateway.New(calculatorService, server, router)
}
