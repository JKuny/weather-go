package model_test

import (
	"github.com/jkuny/weather-go/internal/apis/open_meteo/model"
	"testing"
)

const (
	earlyTime         = "2006-01-02T09:04"
	expectedEarlyTime = "09:04AM"
	lateTime          = "2006-01-02T22:04"
	expectedLateTime  = "10:04PM"
	invalidTime       = "2006-01-02T15:04 GMT"
	expectedError     = "Expected error, got nil"
)

func TestEarlyTimeFormat(t *testing.T) {
	got := model.ConvertTime(earlyTime)
	if got != expectedEarlyTime {
		t.Errorf("Got %s, wanted %s", got, expectedEarlyTime)
	}
}

func TestLateTimeFormat(t *testing.T) {
	got := model.ConvertTime(lateTime)
	if got != expectedLateTime {
		t.Errorf("Got %s, wanted %s", got, expectedLateTime)
	}
}

func TestInvalidTimeFormat(t *testing.T) {
	got := model.ConvertTime(invalidTime)
	if got != "" {
		t.Errorf(expectedError)
	}
}
