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
	_sumOperation      = "sum"
	_subtractOperation = "subtraction"
	_multiplyOperation = "multiplication"
	_divideOperation   = "division"
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
func (s *experimenterService) processResult(op *operator) *evaluator {
	ev := newEvaluator(op)
	s.doExperiment(op, ev)
	return ev
}

// doExperiment perform the interactions of the experiment.
func (s *experimenterService) doExperiment(op *operator, ev *evaluator) {
	log.Println(fmt.Sprintf("Running for %v", op))
	rs := newResultsStore(op)
	for baseIndex := 0; baseIndex < op.Interactions; baseIndex = baseIndex + op.BatchSize {
		s.doSequentialRequests(op, rs, baseIndex)
		time.Sleep(time.Duration(op.IntervalBetweenBatchesInMilliseconds) * time.Millisecond)
	}
	ev.RTTSInMicroseconds = rs.getRTTs()
	ev.Failures = rs.getFailures()
}

func (s *experimenterService) doSequentialRequests(op *operator, resultsStore *resultsStore, baseIndex int) {
	ch := make(chan bool, op.BatchSize)
	for batchInternalIndex := 0; batchInternalIndex < op.BatchSize; batchInternalIndex++ {
		actualIndex := baseIndex + batchInternalIndex
		go func(actualIndex int) {
			rtt, err := s.doOperation(actualIndex, op)
			resultsStore.setRTTAt(actualIndex, rtt.Microseconds())
			if err != nil {
				log.Printf("Faiulure in request %d reason: %s\n", actualIndex, err.Error())
				resultsStore.addFailure()
			}
			ch <- true
		}(actualIndex)
	}
	// Waiting for every routine to finish
	waitGoRoutines(ch, op.BatchSize)
}

func waitGoRoutines(ch chan bool, waitSize int) {
	for index := 0; index < waitSize; index++ {
		<-ch
	}
}

// doOperation performs the operation based on an index.
func (s *experimenterService) doOperation(index int, operator *operator) (time.Duration, error) {
	numberOfOperations := len(operator.AllowedOperations)
	remainder := index % numberOfOperations
	operation := operator.AllowedOperations[remainder]
	requestData := map[string]any{
		"firstNumber":  rand.Float64() * 100,
		"secondNumber": (rand.Float64() * 100) + 1,
	}
	var requestFunction func(map[string]any) (map[string]any, error, time.Duration)
	switch operation {
	case _sumOperation:
		requestFunction = s.requester.RequestAdd
	case _subtractOperation:
		requestFunction = s.requester.RequestSubtract
	case _multiplyOperation:
		requestFunction = s.requester.RequestMultiply
	case _divideOperation:
		requestFunction = s.requester.RequestDivide
	}
	_, err, rtt := requestFunction(requestData)
	return rtt, err
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
