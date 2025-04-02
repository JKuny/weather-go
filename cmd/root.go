package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"weather-go/internal/apis/open_meteo"
)

var rootCmd = &cobra.Command{
	Use:   "weather-go",
	Short: "weather-go is a small CLI weather application",
	Long: `A CLI weather application in Go. Allows for displaying the current weather
in multiple locations via multiple APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting your configured weather...")
		fmt.Println(open_meteo.GetCurrentWeather())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
