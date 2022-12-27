package experimenter

import "net/http"

//go:generate moq -out mocks/service_mock.go -pkg mocks . Service

type Service interface {
	Experiment(http.ResponseWriter, *http.Request)
}
