package model

type WeatherCondition struct {
	Code        int
	Description string
}

// WeatherConditionsMap implementation of the following table: https://open-meteo.com/en/docs#weather_variable_documentation
var WeatherConditionsMap = map[int]WeatherCondition{
	0: {Code: 0, Description: "Clear Sky"},
	1: {Code: 1, Description: "Mainly Clear"},
	2: {Code: 2, Description: "Partly Cloudy"},
	3: {Code: 3, Description: "Overcast"},

	45: {Code: 45, Description: "Fog"},
	48: {Code: 48, Description: "Depositing Rime Fog"},

	51: {Code: 51, Description: "Drizzle: Light"},
	53: {Code: 53, Description: "Drizzle: Moderate"},
	55: {Code: 55, Description: "Drizzle: Dense"},

	56: {Code: 56, Description: "Freezing Drizzle: Light"},
	57: {Code: 57, Description: "Freezing Drizzle: Dense"},

	61: {Code: 61, Description: "Rain: Light"},
	63: {Code: 63, Description: "Rain: Moderate"},
	65: {Code: 65, Description: "Rain: Dense"},

	66: {Code: 66, Description: "Freezing Rain: Light"},
	67: {Code: 67, Description: "Freezing Rain: Dense"},

	71: {Code: 71, Description: "Snowfall: Light"},
	73: {Code: 73, Description: "Snowfall: Moderate"},
	75: {Code: 75, Description: "Snowfall: Heavy"},

	77: {Code: 77, Description: "Snow Grains"},

	80: {Code: 80, Description: "Rain Showers: Light"},
	81: {Code: 81, Description: "Rain Showers: Moderate"},
	82: {Code: 82, Description: "Rain Showers: Violent"},

	85: {Code: 85, Description: "Snow Showers: Light"},
	86: {Code: 86, Description: "Snow Showers: Heavy"},

	95: {Code: 95, Description: "Thunderstorm: Slight or Moderate"},

	96: {Code: 96, Description: "Thunderstorm with slight hail"},
	99: {Code: 99, Description: "Thunderstorm with heavy hail"},
}
