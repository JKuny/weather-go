package display_test

import (
	"github.com/jkuny/weather-go/internal/display"
	"testing"
)

const (
	earlyTime         = "2006-01-02T09:04"
	expectedEarlyTime = "09:04AM"
	lateTime          = "2006-01-02T22:04"
	expectedLateTime  = "10:04PM"
	invalidTime       = "2006-01-02T15:04 GMT"
	expectedDate      = "2006-01-02"
	expectedError     = "Expected error, got nil"
)

func TestEarlyTimeFormat(t *testing.T) {
	got := display.FormatTime(earlyTime)
	if got != expectedEarlyTime {
		t.Errorf("Got %s, wanted %s", got, expectedEarlyTime)
	}
}

func TestLateTimeFormat(t *testing.T) {
	got := display.FormatTime(lateTime)
	if got != expectedLateTime {
		t.Errorf("Got %s, wanted %s", got, expectedLateTime)
	}
}

func TestInvalidTimeFormat(t *testing.T) {
	got := display.FormatTime(invalidTime)
	if got != "" {
		t.Errorf(expectedError)
	}
}

func TestFormatDate(t *testing.T) {
	got := display.FormatDate(earlyTime)
	if got != expectedDate {
		t.Errorf(expectedError)
	}
}
