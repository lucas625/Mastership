package service

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/lucas625/Mastership/experimenter-service/experimenter"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/calculator_requester"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/errors"
)

const (
	_addOperation = iota
	_subtractOperation
	_multiplyOperation
	_divideOperation
)

// An implementation of the experimenter.Service interface.
type experimenterService struct {
	requester calculator_requester.CalculatorRequester
}

var _ experimenter.Service = &experimenterService{}

// New creates a new experimenter service.
func New(requester calculator_requester.CalculatorRequester) experimenter.Service {
	return &experimenterService{requester: requester}
}

// Experiment gets a json with parameters described on the operator struct and runs the experiment.
func (s *experimenterService) Experiment(responseWriter http.ResponseWriter, request *http.Request) {
	s.processRequest(responseWriter, request)
}

// processRequest processes the request.
func (s *experimenterService) processRequest(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	if request.Method == http.MethodOptions {
		return
	}
	operator, err := s.parseRequest(responseWriter, request)
	if err == nil {
		result := s.processResult(operator)
		log.Println("Finished evaluation")
		log.Println("Processing results")
		err := result.ProcessResults()
		if err != nil {
			log.Printf("Failed to process results due to: %s", err.Error())
			http.Error(responseWriter, err.Error(), 500)
			return
		}
		s.sendResponse(responseWriter, result)
	}
}

// parseRequest parses the request.
func (s *experimenterService) parseRequest(responseWriter http.ResponseWriter, request *http.Request) (*operator, error) {
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
		case errors.ValidationError:
			http.Error(responseWriter, err.Error(), 400)
		default:
			http.Error(responseWriter, err.Error(), 500)
		}
		log.Println(err)
		return nil, err
	}
	return operator, nil
}

// processResult calculates the result.
func (s *experimenterService) processResult(operator *operator) *evaluator {
	evaluator := newEvaluator(operator)
	s.doExperiment(operator, evaluator)
	return evaluator
}

// doExperiment perform the interactions of the experiment.
func (s *experimenterService) doExperiment(operator *operator, evaluator *evaluator) {
	log.Println(fmt.Sprintf("Running for %v", operator))
	for baseIndex := 0; baseIndex < operator.Interactions; baseIndex = baseIndex + operator.BatchSize {
		for batchInternalIndex := 0; batchInternalIndex < operator.BatchSize; batchInternalIndex++ {
			actualIndex := baseIndex + batchInternalIndex
			startTime := time.Now()
			err := s.doOperation(actualIndex)
			rtt := time.Now().Sub(startTime)
			if err != nil {
				log.Printf("Faiulure in request %d reason: %s\n", actualIndex, err.Error())
				evaluator.Failures++
			}
			evaluator.RTTSInMicroseconds[actualIndex] = rtt.Microseconds()
		}
		time.Sleep(time.Duration(operator.IntervalBetweenBatchesInMilliseconds) * time.Millisecond)
	}
}

// doOperation performs the operation based on an index.
func (s *experimenterService) doOperation(index int) error {
	numberOfOperations := 4
	remainder := index % numberOfOperations
	requestData := map[string]any{
		"firstNumber":  rand.Float64() * 100,
		"secondNumber": (rand.Float64() * 100) + 1,
	}
	var err error = nil
	switch remainder {
	case _addOperation:
		_, err = s.requester.RequestAdd(requestData)
	case _subtractOperation:
		_, err = s.requester.RequestSubtract(requestData)
	case _multiplyOperation:
		_, err = s.requester.RequestMultiply(requestData)
	case _divideOperation:
		_, err = s.requester.RequestDivide(requestData)
	}
	return err
}

// sendResponse sends the response.
func (s *experimenterService) sendResponse(responseWriter http.ResponseWriter, result *evaluator) {
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
