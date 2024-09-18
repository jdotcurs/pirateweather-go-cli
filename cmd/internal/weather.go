package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/geocoding"
	"github.com/jdotcurs/pirateweather-go/pkg/models"
	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/jdotcurs/pirateweather-go/pkg/utils"
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
	ForecastCmd.Flags().StringP("location", "l", "", "Location (address or coordinates)")
	ForecastCmd.Flags().StringP("units", "u", "si", "Units to use (si, us, uk, ca)")

	TimeMachineCmd.Flags().StringP("location", "l", "", "Location (address or coordinates)")
	TimeMachineCmd.Flags().StringP("units", "u", "si", "Units to use (si, us, uk, ca)")
	TimeMachineCmd.Flags().StringP("time", "t", "", "Time for historical data (format: YYYY-MM-DD)")
}

func RunInteractiveForecast() {
	location := promptForLocation()
	units := promptForUnits()

	latitude, longitude := getCoordinates(location)
	if latitude == 0 && longitude == 0 {
		fmt.Println("Failed to get coordinates. Exiting.")
		return
	}

	client := getClient()
	forecast, err := client.Forecast(latitude, longitude, pirateweather.WithUnits(units))
	if err != nil {
		fmt.Printf("Error fetching forecast: %v\n", err)
		return
	}

	printForecast(forecast)
}

func RunInteractiveTimeMachine() {
	location := promptForLocation()
	units := promptForUnits()
	timestamp := promptForTimestamp()

	latitude, longitude := getCoordinates(location)
	if latitude == 0 && longitude == 0 {
		fmt.Println("Failed to get coordinates. Exiting.")
		return
	}

	client := getClient()
	forecast, err := client.TimeMachine(latitude, longitude, timestamp, pirateweather.WithUnits(units))
	if err != nil {
		fmt.Printf("Error fetching historical data: %v\n", err)
		return
	}

	printForecast(forecast)
}

func promptForLocation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter location (address, city, or latitude,longitude): ")
	location, _ := reader.ReadString('\n')
	return strings.TrimSpace(location)
}

func promptForUnits() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter units (si, us, uk, ca) [default: si]: ")
	units, _ := reader.ReadString('\n')
	units = strings.TrimSpace(units)
	if units == "" {
		units = "si"
	}
	return units
}

func promptForTimestamp() time.Time {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date for historical data (YYYY-MM-DD): ")
	dateStr, _ := reader.ReadString('\n')
	dateStr = strings.TrimSpace(dateStr)
	timestamp, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Printf("Invalid date format. Using current date.\n")
		return time.Now()
	}
	return timestamp
}

func getCoordinates(location string) (float64, float64) {
	parts := strings.Split(location, ",")
	if len(parts) == 2 {
		lat, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		lon, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if err1 == nil && err2 == nil {
			return lat, lon
		}
	}

	// If not coordinates, use geocoding
	geocodeResult, err := geocoding.ForwardGeocode(location)
	if err != nil {
		fmt.Printf("Error geocoding location: %v\n", err)
		fmt.Println("Please try entering the coordinates directly (latitude,longitude)")
		return 0, 0
	}

	lat, err := strconv.ParseFloat(geocodeResult.Lat, 64)
	if err != nil {
		fmt.Printf("Error parsing latitude: %v\n", err)
		return 0, 0
	}

	lon, err := strconv.ParseFloat(geocodeResult.Lon, 64)
	if err != nil {
		fmt.Printf("Error parsing longitude: %v\n", err)
		return 0, 0
	}

	return lat, lon
}

func getClient() *pirateweather.Client {
	apiKey := os.Getenv("PIRATE_WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: PIRATE_WEATHER_API_KEY environment variable is not set")
		os.Exit(1)
	}
	return pirateweather.NewClient(apiKey)
}

func printForecast(forecast *models.ForecastResponse) {
	fmt.Printf("Location: %.4f, %.4f\n", forecast.Latitude, forecast.Longitude)

	// Get address information
	geocodeResult, err := geocoding.ReverseGeocode(forecast.Latitude, forecast.Longitude)
	if err != nil {
		fmt.Printf("Error getting address information: %v\n", err)
	} else {
		fmt.Printf("Address: %s\n", geocodeResult.DisplayName)
		fmt.Printf("City: %s\n", geocodeResult.Address.City)
		fmt.Printf("State: %s\n", geocodeResult.Address.State)
		fmt.Printf("Country: %s\n", geocodeResult.Address.Country)
	}

	fmt.Printf("Timezone: %s\n", forecast.Timezone)
	fmt.Printf("Time: %s\n", utils.FormatTime(forecast.Currently.Time))
	fmt.Printf("Temperature: %.2f°C\n", forecast.Currently.Temperature)
	fmt.Printf("Feels like: %.2f°C\n", forecast.Currently.ApparentTemperature)
	fmt.Printf("Humidity: %.2f%%\n", forecast.Currently.Humidity*100)
	fmt.Printf("Wind Speed: %.2f m/s\n", forecast.Currently.WindSpeed)
	fmt.Printf("Wind Direction: %.2f°\n", forecast.Currently.WindBearing)
	fmt.Printf("Cloud Cover: %.2f%%\n", forecast.Currently.CloudCover*100)
	fmt.Printf("UV Index: %.1f\n", forecast.Currently.UVIndex)
	fmt.Printf("Visibility: %.2f km\n", forecast.Currently.Visibility)
	fmt.Printf("Fire Index: %.2f\n", forecast.Currently.FireIndex)
	fmt.Printf("Smoke: %.2f\n", forecast.Currently.Smoke)

	if len(forecast.Alerts) > 0 {
		fmt.Println("\nWeather Alerts:")
		for _, alert := range forecast.Alerts {
			fmt.Printf("- %s: %s\n", alert.Title, alert.Description)
		}
	}

	if forecast.Hourly != nil && len(forecast.Hourly.Data) > 0 {
		fmt.Printf("\nHourly forecast available for the next %d hours\n", len(forecast.Hourly.Data))
	}

	if forecast.Daily != nil && len(forecast.Daily.Data) > 0 {
		fmt.Printf("\nDaily forecast available for the next %d days\n", len(forecast.Daily.Data))
	}
}

func runForecast(cmd *cobra.Command, args []string) {
	location, _ := cmd.Flags().GetString("location")
	units, _ := cmd.Flags().GetString("units")

	latitude, longitude := getCoordinates(location)
	if latitude == 0 && longitude == 0 {
		fmt.Println("Failed to get coordinates. Exiting.")
		return
	}

	client := getClient()
	forecast, err := client.Forecast(latitude, longitude, pirateweather.WithUnits(units))
	if err != nil {
		fmt.Printf("Error fetching forecast: %v\n", err)
		return
	}

	printForecast(forecast)
}

func runTimeMachine(cmd *cobra.Command, args []string) {
	location, _ := cmd.Flags().GetString("location")
	units, _ := cmd.Flags().GetString("units")
	timeStr, _ := cmd.Flags().GetString("time")

	latitude, longitude := getCoordinates(location)
	if latitude == 0 && longitude == 0 {
		fmt.Println("Failed to get coordinates. Exiting.")
		return
	}

	timestamp, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		fmt.Printf("Invalid time format: %v\n", err)
		return
	}

	client := getClient()
	forecast, err := client.TimeMachine(latitude, longitude, timestamp, pirateweather.WithUnits(units))
	if err != nil {
		fmt.Printf("Error fetching historical data: %v\n", err)
		return
	}

	printForecast(forecast)
}
