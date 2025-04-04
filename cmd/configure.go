// Package cmd
/*
Copyright © 2025 James Kuny <james.kuny@yahoo.com>
*/
package cmd

import (
	"log"

	"github.com/jkuny/weather-go/internal/display"
	"github.com/jkuny/weather-go/internal/validation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var printConfig = &cobra.Command{
	Use:   "print",
	Short: "Print the current configuration",
	Long:  `Print the current configuration file values in a human readable format.`,
	Run: func(cmd *cobra.Command, args []string) {
		settings := viper.AllSettings()
		keys := make([]string, 0, len(settings))
		for k := range settings {
			keys = append(keys, k)
		}
		display.PrintMap(settings, "")
	},
}

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure functionality of weather-go",
	Long: `Configure functionality of weather-go like:

Default latitude/longitude to get weather for.
Weather website source to use.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		latitude, _ := cmd.Flags().GetString("latitude")
		longitude, _ := cmd.Flags().GetString("longitude")

		// Validate the values passed in to be valid latitude/longitude
		_, err := validation.ValidateLatLong(latitude, longitude)
		if err != nil {
			log.Fatalln(err)
		}

		// Update the config file's values
		viper.Set("app.default-location.latitude", latitude)
		viper.Set("app.default-location.longitude", longitude)
		err = viper.WriteConfig()
		if err != nil {
			log.Fatalln("Issue saving the configuration file:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(printConfig)

	configureCmd.Flags().StringP("latitude", "l", "", "set the default latitude")
	configureCmd.Flags().StringP("longitude", "g", "", "set the default longitude")
	configureCmd.MarkFlagsRequiredTogether("latitude", "longitude")

	// Bind the default-location values to the flags
	err := viper.BindPFlag("app.default-location.latitude", configureCmd.Flags().Lookup("latitude"))
	if err != nil {
		log.Fatalln(err)
	}

	err = viper.BindPFlag("app.default-location.latitude", configureCmd.Flags().Lookup("latitude"))
	if err != nil {
		log.Fatalln(err)
	}
}
