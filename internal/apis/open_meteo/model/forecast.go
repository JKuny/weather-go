package model

import (
	"fmt"
	"github.com/jkuny/weather-go/internal/display"
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
	fmt.Printf("%s -> %s\n", display.FormatDate(forecast.Hourly.Time[0]), display.FormatDate(forecast.Hourly.Time[len(forecast.Hourly.Time)-1]))
	for i := range forecast.Hourly.Time {
		fmt.Printf("     %v --- %.1f %s --- %v%% \n",
			display.FormatTime(forecast.Hourly.Time[i])+" "+forecast.Timezone,
			forecast.Hourly.Temperature[i],
			forecast.HourlyUnits.TemperatureUnits,
			forecast.Hourly.PrecipitationProbability[i],
		)
	}
}
