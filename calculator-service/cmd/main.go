package main

import (
	"log"

	"github.com/lucas625/Mastership/calculator-service/calculator/service_gateway"
)

func main() {
	if err := service_gateway.New().CreateServer(); err != nil {
		log.Fatal(err)
	}
}
