package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucas625/Mastership/experimenter-service/experimenter"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/errors"
)

// An implementation of the experimenter.Service interface.
type experimenterService struct{}

var _ experimenter.Service = experimenterService{}

// New creates a new experimenter service.
func New() experimenter.Service {
	return &experimenterService{}
}

// Experiment gets a json with parameters described on the operator struct and runs the experiment.
func (s experimenterService) Experiment(responseWriter http.ResponseWriter, request *http.Request) {
	processRequest(responseWriter, request)
}

// processRequest processes the request.
func processRequest(responseWriter http.ResponseWriter, request *http.Request) {
	operator, err := parseRequest(responseWriter, request)
	if err == nil {
		result := processResult(operator)
		log.Println("Finished evaluation")
		sendResponse(responseWriter, result)
	}
}

// parseRequest parses the request.
func parseRequest(responseWriter http.ResponseWriter, request *http.Request) (*operator, error) {
	operator, err := requestToOperator(request)
	if err != nil {
		switch err.(type) {
		case errors.UnmarshalError:
			http.Error(responseWriter, err.Error(), 400)
		case errors.KeyNotFoundError:
			http.Error(responseWriter, err.Error(), 400)
		case errors.DecodeError:
			http.Error(responseWriter, err.Error(), 400)
		case errors.InvalidValueError:
			http.Error(responseWriter, err.Error(), 400)
		default:
			http.Error(responseWriter, err.Error(), 500)
		}
		log.Println(err)
		return nil, err
	}
	// TODO: validate request
	return operator, nil
}

// processResult calculates the result.
func processResult(operator *operator) *evaluator {
	evaluator := newEvaluator(operator.interactions)
	doExperiment(operator, evaluator)
	return evaluator
}

func doExperiment(operator *operator, evaluator *evaluator) {
	for index := 0; index < operator.interactions; index = index + operator.batchSize {
		for batchInternalIndex := 0; batchInternalIndex < operator.batchSize; batchInternalIndex++ {
			// Start timer
			// Do request
			// Get time diff
			// Set time diff (in MS) to evaluator.RTTSInMS[index + batchInternalIndex]
		}
	}
}

// sendResponse sends the response.
func sendResponse(responseWriter http.ResponseWriter, result *evaluator) {
	responseAsBytes, err := json.Marshal(result)
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
		log.Println(err.Error())
		return
	}
	_, err = responseWriter.Write(responseAsBytes)
	if err != nil {
		http.Error(responseWriter, err.Error(), 500)
		log.Println(err.Error())
		return
	}
}
