package open_meteo

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
