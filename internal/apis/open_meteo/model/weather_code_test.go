package model_test

import (
	"github.com/jkuny/weather-go/internal/apis/open_meteo/model"
	"testing"
)

const expectedError = "Expected error, got nil"

var expectedCodes = [...]int{
	0, 1, 2, 3, 45, 48, 51, 53, 55, 56, 57, 61, 63, 65, 66, 67, 71, 73, 75, 77, 80, 81, 82, 85, 86, 95, 96, 99,
}

func TestValidWeatherCodes(t *testing.T) {
	for i := range expectedCodes {
		got := model.WeatherConditionsMap[expectedCodes[i]].Description
		if got == "" {
			t.Errorf(expectedError)
		}
	}
}

// TestInvalidWeatherCodes a code that isn't in the list should return nothing
func TestInvalidWeatherCodes(t *testing.T) {
	got := model.WeatherConditionsMap[1000].Description
	if got != "" {
		t.Errorf(expectedError)
	}
}
