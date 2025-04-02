package open_weather

import (
	"net/http"
)

func GetCurrentWeather() string {
	resp, err := http.Get("https://jameskuny.com/")
	if err != nil {
		return ""
	}

	return resp.Status
}
