package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pirateweather",
	Short: "A CLI for the Pirate Weather API",
	Long:  `PirateWeather CLI provides a command-line interface to fetch weather data using the Pirate Weather API.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// rootCmd.AddCommand(weather.ForecastCmd)
	// rootCmd.AddCommand(weather.TimeMachineCmd)
}
