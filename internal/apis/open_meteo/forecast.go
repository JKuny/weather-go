package open_meteo

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
		Time             string `json:"time"`
		TemperatureUnits string `json:"temperature_2m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time        []string  `json:"time"`
		Temperature []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

// DisplayForecast displays a display-friendly version of the Forecast struct
func DisplayForecast(forecast Forecast, numberOfDays string) {
	fmt.Printf("Elevation: %v \n", forecast.Elevation)
	fmt.Printf("Forecast for the next %s day(s):\n", numberOfDays)
	for i := range forecast.Hourly.Time {
		fmt.Printf("     %v --- %v --- %.1f %s\n",
			convertTime(forecast.Hourly.Time[i]+" "+forecast.TimezoneAbbrev),
			forecast.Hourly.,
			forecast.Hourly.Temperature[i],
			forecast.HourlyUnits.TemperatureUnits)
	}
}

func convertTime(timeStr string) string {
	layout := "2006-01-02T15:04 MST"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return ""
	}

	estLocation, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("Error loading EST timezone: %v\n", err)
		return ""
	}

	return parsedTime.In(estLocation).Format(layout)
}
