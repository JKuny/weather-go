// Package cmd
/*
Copyright © 2025 James Kuny <james.kuny@yahoo.com>
*/

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jkuny/weather-go/internal/apis/open_meteo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "weather-go",
	Short: "Retrieve the immediate forecast for the default latitude/longitude configured with the application",
	Long: `A CLI weather application in Go. Allows for displaying the current weather
in multiple locations via multiple APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultLatitude := viper.GetString("app.default-location.latitude")
		defaultLongitude := viper.GetString("app.default-location.longitude")

		if defaultLatitude == "" || defaultLongitude == "" {
			log.Fatal("Latitude/Longitude not configured.")
		}

		fmt.Printf("Getting weather for %s/%s...\n", defaultLatitude, defaultLongitude)
		body, err := open_meteo.GetCurrentWeather(defaultLatitude, defaultLongitude)
		if err != nil {
			log.Fatalf("Error while getting weather from OpenMeteo: %s", err)
		}

		formattedData, err := open_meteo.ParseData(body)
		if err != nil {
			log.Fatalf("Error while parsing weather data: %s", err)
		}

		fmt.Println(formattedData)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
