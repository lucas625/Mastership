package configuration

import (
	"log"

	env "github.com/caarlos0/env/v6"
)

type Configuration struct {
	Port int `env:"MSC_CALCULATOR_SERVICE_PORT" envDefault:"8001"`
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
