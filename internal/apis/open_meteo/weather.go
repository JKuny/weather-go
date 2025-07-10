package open_meteo

import (
	"fmt"
	"time"
)

type Weather struct {
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

// DisplayWeather displays a display-friendly version of the Weather struct
func DisplayWeather(weather Weather) {
	fmt.Printf("Latitude: %v, %v\n", weather.Latitude, weather.Longitude)

	fmt.Println("Temperature:")
	for i := 0; i < len(weather.Hourly.Time); i++ {
		fmt.Printf("     %v     ---     %v %s\n", convertTime(weather.Hourly.Time[i]+" "+weather.TimezoneAbbrev), weather.Hourly.Temperature[i], weather.HourlyUnits.TemperatureUnits)
	}
	fmt.Printf("Elevation: %v \n", weather.Elevation)
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
