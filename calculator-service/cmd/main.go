package main

import (
	"log"

	"github.com/lucas625/Mastership/calculator-service/calculator"
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
	configuration, err := calculator.NewConfiguration()
	if err != nil {
		return err
	}
	calculatorService := service.New()
	gateway := service_gateway.New(calculatorService, configuration)
	return gateway.CreateServer()
}
