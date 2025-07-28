package open_meteo_test

import (
	"testing"
)

const (
	expectedError = "Expected error, got nil"
)

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Errorf(expectedError)
	}
}
