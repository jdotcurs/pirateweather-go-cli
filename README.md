# PirateWeather Go CLI

A command-line interface for the Pirate Weather API, built using the Pirate Weather Go SDK.

## Installation

To install the Pirate Weather CLI, use the following command:

```bash
go install github.com/jdotcurs/pirateweather-go-cli@latest
```

## Usage

Before using the CLI, set your Pirate Weather API key as an environment variable:

```bash
export PIRATE_WEATHER_API_KEY=<your-api-key>
```

You can find your API key in your [Pirate Weather account settings](https://pirate-weather.apiable.io).


### Interactive Mode

Run the CLI without arguments to enter interactive mode:

```bash
pirateweather
```

This will guide you through the process of getting weather information.

### Command-line Mode

#### Get current forecast

```bash
pirateweather forecast --location "New York" --units si
```

This will display the current forecast for the specified location in SI units (Celsius, km/h, hPa).

#### Get current forecast for a specific location

```bash
pirateweather forecast --location "New York" --units si
```


#### Get historical weather data

```bash
pirateweather timemachine --location "40.7128,-74.0060" --units si --time 2023-05-01
```


## Options

- `--location, -l`: Location (address, city, or latitude,longitude) (required)
- `--units, -u`: Units to use (si, us, uk, ca) (default: si)
- `--time, -t`: Time for historical data (format: YYYY-MM-DD) (required for timemachine command)

## Features

- Interactive mode for guided weather queries
- Support for both current forecasts and historical weather data
- Flexible location input (address, city name, or coordinates)
- Multiple unit system support (SI, US, UK, Canada)
- Detailed weather information including temperature, humidity, wind speed, and more
- Error handling and user-friendly messages

## Examples

1. Get current weather for London in SI units:

```bash
pirateweather forecast --location "London" --units si
```

2. Get historical weather for New York on January 1, 2023, in US units:

```bash
pirateweather timemachine --location "New York" --units us --time 2023-01-01
```

3. Get current weather for specific coordinates in UK units:

```bash
pirateweather forecast --location "51.5074,-0.1278" --units uk
```

## Contributing

Contributions to the Pirate Weather CLI are welcome! Please follow these steps:

1. Fork the repository
2. Create a new branch for your feature
3. Write tests for your new feature
4. Implement your feature
5. Run the test suite
6. Create a pull request

Please ensure your code adheres to the existing style and passes all tests.

## License

This project is licensed under the MIT License.
