package main

import (
	"log"

	"github.com/spf13/viper"
	"weather-go/cmd"
)

type Configuration struct {
	App struct {
		Name    string
		Version string
		Source  string
	}
}

func main() {
	// Viper configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Run
	cmd.Execute()
}
