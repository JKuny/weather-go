// Package open_weather
/*
Copyright © 2025 James Kuny <james.kuny@yahoo.com>
*/
package open_weather

import (
	"net/http"
)

var baseUrl = "https://jameskuny.com/"

func GetCurrentWeather() string {
	resp, err := http.Get(baseUrl)
	if err != nil {
		return ""
	}

	return resp.Status
}
