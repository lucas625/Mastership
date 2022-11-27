package configuration

import (
	"log"

	env "github.com/caarlos0/env/v6"
)

type Configuration struct {
	CalculatorServiceAddress string `env:"MSC_CALCULATOR_SERVICE_ADDRESS" envDefault:"http://127.0.0.1:8000"`
	Port                     int    `env:"MSC_EXPERIMENTER_SERVICE_PORT" envDefault:"8001"`
}

func New() (*Configuration, error) {
	log.Printf("Reading environment variables")
	configuration := &Configuration{}
	err := env.Parse(configuration)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}
