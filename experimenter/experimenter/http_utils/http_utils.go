package http_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lucas625/Mastership/experimenter-service/experimenter/errors"
	"io"
	"io/ioutil"
	"net/http"
)

// PostRequest performs a post request.
func PostRequest(url string, data map[string]any) (map[string]any, error) {
	ioReaderData, err := dataToRequestData(data)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/json", ioReaderData)
	if err != nil {
		return nil, err
	}

	return getResponseData(response)
}

// dataToRequestData parses the map[string]any to the io.Reader expected by the http library.
func dataToRequestData(data map[string]any) (io.Reader, error) {
	ioReaderData := new(bytes.Buffer)
	err := json.NewEncoder(ioReaderData).Encode(data)
	if err != nil {
		return nil, err
	}
	return ioReaderData, nil
}

// getResponseData gets the response data from the http.Response as a map[string]any.
func getResponseData(response *http.Response) (map[string]any, error) {
	if response.StatusCode != http.StatusOK {
		return nil, errors.NewRequestError(fmt.Sprintf("Request status is %d", response.StatusCode))
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var parsedResponseData map[string]any
	// TODO: The response is arriving as a number, change to return a json
	err = json.Unmarshal(responseData, &parsedResponseData)
	return parsedResponseData, err
}
