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
	temperature          string = "temperature_2m"
	precipitation        string = "precipitation_probability"
	precipitation_amount string = "precipitation"
	weather_code         string = "weather_code"
	wind_speed           string = "wind_speed_10m"
)

var baseUrl = "https://api.open-meteo.com/v1/forecast"

// GetCurrentWeather Gets the current weather forecast for the next two minutes
func GetCurrentWeather(latitude string, longitude string) (string, error) {
	client := &http.Client{}
	endpoint, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Encode a set of parameters into the URL
	endpoint.RawQuery = url.Values{
		"latitude":  {latitude},
		"longitude": {longitude},
		"hourly": {
			temperature,
			precipitation,
			precipitation_amount,
			weather_code,
			wind_speed,
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
func ParseData(body string) (Weather, error) {
	var weather Weather
	err := json.Unmarshal([]byte(body), &weather)
	if err != nil {
		log.Fatalf("Issue parsing JSON: %s", err)
		return Weather{}, err
	}
	return weather, nil
}
