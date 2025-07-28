package model

import (
	"fmt"
	"time"
)

type Forecast struct {
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	GenerationTimeMs float64 `json:"generationtime_ms"`
	UTCOffsetSeconds int     `json:"utc_offset_seconds"`
	Timezone         string  `json:"timezone"`
	TimezoneAbbrev   string  `json:"timezone_abbreviation"`
	Elevation        float64 `json:"elevation"`
	HourlyUnits      struct {
		Time                     string `json:"time"`
		TemperatureUnits         string `json:"temperature_2m"`
		PrecipitationProbability string `json:"precipitation_probability"`
	} `json:"hourly_units"`
	Hourly struct {
		Time                     []string  `json:"time"`
		Temperature              []float64 `json:"temperature_2m"`
		PrecipitationProbability []float64 `json:"precipitation_probability"`
	} `json:"hourly"`
}

// DisplayForecast displays a display-friendly version of the Forecast struct
func DisplayForecast(forecast Forecast, numberOfDays string) {
	fmt.Printf("Elevation: %v \n", forecast.Elevation)
	fmt.Printf("Forecast for the next %s day(s):\n", numberOfDays)
	for i := range forecast.Hourly.Time {
		fmt.Printf("     %v --- %.1f %s --- %v%% \n",
			ConvertTime(forecast.Hourly.Time[i])+" "+forecast.Timezone,
			forecast.Hourly.Temperature[i],
			forecast.HourlyUnits.TemperatureUnits,
			forecast.Hourly.PrecipitationProbability[i],
		)
	}
}

// ConvertTime to change the time format to something more readable
func ConvertTime(timeStr string) string {
	layout := "2006-01-02T15:04"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return ""
	}

	// Prepend a 0 if the time is =6 to keep things uniform
	kitchen := parsedTime.Format(time.Kitchen)
	if len(kitchen) == 6 {
		return "0" + kitchen
	} else {
		return kitchen
	}
}
