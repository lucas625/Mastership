package service_gateway

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/lucas625/Mastership/calculator-service/calculator"
	"github.com/lucas625/Mastership/calculator-service/calculator/service"
)

type restGateway struct{}

func New() calculator.ServiceGateway {
	return restGateway{}
}

func (restGateway) CreateServer() error {
	calculatorService := service.New()
	// TODO: split into different functions, add tests, add configuration for env
	router := mux.NewRouter()
	router.HandleFunc("/add", calculatorService.Add)
	router.HandleFunc("/divide", calculatorService.Divide)
	router.HandleFunc("/multiply", calculatorService.Multiply)
	router.HandleFunc("/subtract", calculatorService.Subtract)

	server := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 1 * time.Minute,
		ReadTimeout:  1 * time.Minute,
	}

	log.Println("Server running!")
	return server.ListenAndServe()
}
