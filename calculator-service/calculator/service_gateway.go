package calculator

//go:generate moq -out service_gateway_mock.go . ServiceGateway

type ServiceGateway interface {
	CreateServer() error
}
