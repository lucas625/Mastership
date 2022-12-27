package configuration

import (
	"testing"

	"github.com/matryer/is"
)

// TestConfiguration_New tests the configuration's new method.
func TestConfiguration_New(t *testing.T) {
	expectedConfiguration := &Configuration{Port: 8002}
	config, err := New()

	verifier := is.New(t)
	verifier.Equal(expectedConfiguration, config)
	verifier.Equal(nil, err)
}
