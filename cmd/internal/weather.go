package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/spf13/cobra"
)

var ForecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get weather forecast for a location",
	Run:   runForecast,
}

var TimeMachineCmd = &cobra.Command{
	Use:   "timemachine",
	Short: "Get historical weather data for a location",
	Run:   runTimeMachine,
}

func init() {
	ForecastCmd.Flags().Float64P("latitude", "lat", 0, "Latitude of the location")
	ForecastCmd.Flags().Float64P("longitude", "lon", 0, "Longitude of the location")
	ForecastCmd.Flags().StringP("units", "u", "si", "Units to use (si, us, uk, ca)")

	TimeMachineCmd.Flags().Float64P("latitude", "lat", 0, "Latitude of the location")
	TimeMachineCmd.Flags().Float64P("longitude", "lon", 0, "Longitude of the location")
	TimeMachineCmd.Flags().StringP("units", "u", "si", "Units to use (si, us, uk, ca)")
	TimeMachineCmd.Flags().StringP("time", "t", "", "Time for historical data (format: YYYY-MM-DD)")
}

func runForecast(cmd *cobra.Command, args []string) {
	latitude, _ := cmd.Flags().GetFloat64("latitude")
	longitude, _ := cmd.Flags().GetFloat64("longitude")
	units, _ := cmd.Flags().GetString("units")

	client := getClient()
	forecast, err := client.Forecast(latitude, longitude, pirateweather.WithUnits(units))
	if err != nil {
		fmt.Printf("Error fetching forecast: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Forecast: %v\n", forecast)
	// printForecast(forecast)
}

func runTimeMachine(cmd *cobra.Command, args []string) {
	latitude, _ := cmd.Flags().GetFloat64("latitude")
	longitude, _ := cmd.Flags().GetFloat64("longitude")
	units, _ := cmd.Flags().GetString("units")
	timeStr, _ := cmd.Flags().GetString("time")

	timestamp, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		fmt.Printf("Invalid time format: %v\n", err)
		os.Exit(1)
	}

	client := getClient()
	forecast, err := client.TimeMachine(latitude, longitude, timestamp, pirateweather.WithUnits(units))
	if err != nil {
		fmt.Printf("Error fetching historical data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Forecast: %v\n", forecast)
	// printForecast(forecast)
}

func getClient() *pirateweather.Client {
	apiKey := os.Getenv("PIRATE_WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: PIRATE_WEATHER_API_KEY environment variable is not set")
		os.Exit(1)
	}
	return pirateweather.NewClient(apiKey)
}

// func printForecast(forecast *pirateweather.ForecastResponse) {
// 	fmt.Printf("Temperature: %.2f°C\n", forecast.Currently.Temperature)
// 	fmt.Printf("Humidity: %.2f%%\n", forecast.Currently.Humidity*100)
// 	fmt.Printf("Wind Speed: %.2f m/s\n", forecast.Currently.WindSpeed)
// }
