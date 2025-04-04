package validation_test

import (
	"testing"

	"github.com/jkuny/weather-go/internal/validation"
)

func TestValidateLatLong(t *testing.T) {
	lat, long := "37.7749", "-122.4194"
	got := validation.ValidateLatLong(lat, long)
	if got != true {
		t.Errorf("Got %t, wanted %t", got, true)
	}
}
