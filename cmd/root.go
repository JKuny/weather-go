package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"weather-go/internal/apis/open_meteo"
)

var rootCmd = &cobra.Command{
	Use:   "weather-go",
	Short: "weather-go is a small CLI weather application",
	Long: `A CLI weather application in Go. Allows for displaying the current weather
in multiple locations via multiple APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultLatitude := viper.GetString("app.default-location.latitude")
		defaultLongitude := viper.GetString("app.default-location.longitude")

		if defaultLatitude == "" || defaultLongitude == "" {
			log.Fatal("Latitude/Longitude not configured.")
		}

		fmt.Printf("Getting weather for %s/%s...\n", defaultLatitude, defaultLongitude)
		fmt.Println(open_meteo.GetCurrentWeather(defaultLatitude, defaultLongitude))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
