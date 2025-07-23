// Package open_meteo
/*
Copyright © 2025 James Kuny <james.kuny@yahoo.com>
*/
package open_meteo

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Declare variables to hold names of data to retrieve
var (
	elevation                = "elevation"
	precipitation            = "precipitation"
	precipitationProbability = "precipitation_probability"
	relative_humidity_2m     = "relative_humidity_2m"
	temperature              = "temperature_2m"
	weatherCode              = "weather_code"
	windSpeed                = "wind_speed_10m"
)

var baseUrl = "https://api.open-meteo.com/v1/forecast"

// GetForecast Gets the current weather forecast for the next two minutes
func GetForecast(latitude string, longitude string, numberOfDays string) (string, error) {
	client := &http.Client{}
	endpoint, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Encode a set of parameters into the URL
	endpoint.RawQuery = url.Values{
		"latitude":           {latitude},
		"longitude":          {longitude},
		"forecast_days":      {numberOfDays},
		"wind_speed_unit":    {"mph"},
		"temperature_unit":   {"fahrenheit"},
		"precipitation_unit": {"inch"},
		"hourly": {
			temperature,
			precipitationProbability,
			precipitation,
			weatherCode,
			windSpeed,
			weatherCode,
			relative_humidity_2m,
		},
	}.Encode()

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(body), nil
}

// ParseData Parses an OpenMeteo JSON return into something more display-friendly.
func ParseData(body string) (Forecast, error) {
	var forecast Forecast
	err := json.Unmarshal([]byte(body), &forecast)
	if err != nil {
		log.Fatalf("Issue parsing JSON: %s", err)
		return Forecast{}, err
	}
	return forecast, nil
}
