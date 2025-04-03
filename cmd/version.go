// Package cmd
/*
Copyright © 2025 James Kuny <james.kuny@yahoo.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number for weather-go",
	Long:  `Print the version number for weather-go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("app.version"))
	},
}
