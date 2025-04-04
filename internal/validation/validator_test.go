package validation_test

import (
	"testing"

	"github.com/jkuny/weather-go/internal/validation"
)

const (
	validLat           = "37.7749"
	validLong          = "-122.4194"
	invalidLat         = "not a number"
	invalidLong        = "not a number"
	outOfBoundsLat     = "91"
	outOfBoundsLong    = "181"
	outOfBoundsLatNeg  = "-91"
	outOfBoundsLongNeg = "-181"
	expectedError      = "Expected error, got nil"
)

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Errorf(expectedError)
	}
}

func TestValidateLatLongSuccess(t *testing.T) {
	got, _ := validation.ValidateLatLong(validLat, validLong)
	if !got {
		t.Errorf("Got %t, wanted %t", got, true)
	}
}

func TestValidateLatLongInvalidInputs(t *testing.T) {

	t.Run("invalid lat", func(t *testing.T) {
		_, err := validation.ValidateLatLong(invalidLat, invalidLong)
		assertError(t, err)
	})

	t.Run("invalid long", func(t *testing.T) {
		_, err := validation.ValidateLatLong(validLat, invalidLong)
		assertError(t, err)
	})

	t.Run("out of bounds", func(t *testing.T) {
		_, err := validation.ValidateLatLong(outOfBoundsLat, outOfBoundsLong)
		assertError(t, err)
	})

	t.Run("out of bounds negative", func(t *testing.T) {
		_, err := validation.ValidateLatLong(outOfBoundsLatNeg, outOfBoundsLongNeg)
		assertError(t, err)
	})
}
