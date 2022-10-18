package errors

// DecodeError happens when trying to decode.
type DecodeError struct{}

// NewDecodeError creates a new DecodeError.
func NewDecodeError() error {
	return DecodeError{}
}

// Error returns the string of the error.
func (err DecodeError) Error() string {
	return "Failed to decode"
}

// UnmarshalError happens when trying to unmarshal.
type UnmarshalError struct{}

// NewUnmarshalError creates a new UnmarshalError.
func NewUnmarshalError() error {
	return UnmarshalError{}
}

// Error returns the string of the error.
func (err UnmarshalError) Error() string {
	return "Failed to unmarshal"
}
