package errors

import "fmt"

type RequestError struct {
	reason string
}

var _ error = RequestError{}

func NewRequestError(reason string) error {
	return &RequestError{reason: reason}
}

func (err RequestError) Error() string {
	return fmt.Sprintf("Request failed with reason: %s", err.reason)
}
