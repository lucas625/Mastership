package experimenter

//go:generate moq -out mocks/service_gateway_mock.go -pkg mocks . ServiceGateway

type ServiceGateway interface {
	Serve() error
}
