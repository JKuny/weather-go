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
	shortPercentage   = 1
	normalPercentage  = 50
	largePercentage   = 100
	expectedError     = "Expected error, got nil"
)

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Errorf(expectedError)
	}
}

func TestFormatDate(t *testing.T) {
	got := display.FormatDate(earlyTime)
	if got != expectedDate {
		t.Errorf(expectedError)
	}
}

func TestTimeFormat(t *testing.T) {
	t.Run("early time", func(t *testing.T) {
		got := display.FormatTime(earlyTime)
		if got != expectedEarlyTime {
			t.Errorf("Got %s, wanted %s", got, expectedEarlyTime)
		}
	})

	t.Run("late time", func(t *testing.T) {
		got := display.FormatTime(lateTime)
		if got != expectedLateTime {
			t.Errorf("Got %s, wanted %s", got, expectedLateTime)
		}
	})

	t.Run("invalid time", func(t *testing.T) {
		got := display.FormatTime(invalidTime)
		if got != "" {
			t.Errorf(expectedError)
		}
	})
}

func TestFormatPercentage(t *testing.T) {
	t.Run("short percentage", func(t *testing.T) {
		got := display.FormatPercentage(shortPercentage)
		if got != "01%" {
			assertError(t, nil)
		}
	})

	t.Run("normal percentage", func(t *testing.T) {
		got := display.FormatPercentage(normalPercentage)
		if got != "50%" {
			assertError(t, nil)
		}
	})

	t.Run("large percentage", func(t *testing.T) {
		got := display.FormatPercentage(largePercentage)
		if got != "100%" {
			assertError(t, nil)
		}
	})
}
