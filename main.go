package main

import (
	"fmt"
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

	appName := viper.GetString("app.name")
	version := viper.GetString("app.version")
	source := viper.GetString("app.source")
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Source: %s\n", source)

	// Run
	cmd.Execute()
}
