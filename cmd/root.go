package cmd

import (
	"fmt"
	"os"

	"github.com/jdotcurs/pirateweather-go-cli/cmd/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pirateweather",
	Short: "A CLI for the Pirate Weather API",
	Long:  `PirateWeather CLI provides a command-line interface to fetch weather data using the Pirate Weather API.`,
	Run:   runInteractiveMode,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(internal.ForecastCmd)
	rootCmd.AddCommand(internal.TimeMachineCmd)
}

func runInteractiveMode(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to the PirateWeather CLI!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Get current forecast")
	fmt.Println("2. Get historical weather data")

	var choice string
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		internal.RunInteractiveForecast()
	case "2":
		internal.RunInteractiveTimeMachine()
	default:
		fmt.Println("Invalid choice. Exiting.")
	}
}
