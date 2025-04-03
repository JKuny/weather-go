// Package open_meteo
/*
Copyright © 2025 James Kuny <james.kuny@yahoo.com>
*/
package open_meteo

import (
	"log"
	"net/http"
	"net/url"
)

var baseUrl = "https://api.open-meteo.com/v1/forecast"

// GetCurrentWeather Gets the current weather forecast for the next two minutes
func GetCurrentWeather(latitude string, longitude string) string {
	client := &http.Client{}
	endpoint, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Encode a set of parameters into the URL
	endpoint.RawQuery = url.Values{
		"latitude":  {latitude},
		"longitude": {longitude},
		"hourly":    {"temperature_2m"},
	}.Encode()

	req, err := http.NewRequest("GET", endpoint.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return resp.Status
}
