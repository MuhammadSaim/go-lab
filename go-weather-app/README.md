# Go Weather App

This Go application fetches real-time weather data for Toronto using the Tomorrow.io weather API and displays key weather metrics such as temperature, humidity, and wind speed.

## Features
- Fetches real-time weather data.
- Parses JSON response into structured Go types.
- Displays temperature, humidity, and wind speed.
- Handles errors gracefully with debugging support.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.23 or later)
- Internet connection to access the API

## Usage

1. Run the application:
   ```sh
   go run .
   ```

2. Sample output:
   ```
   Weather Data for: Old Toronto, Toronto, Golden Horseshoe, Ontario, Canada
   Temperature: 1.88 °C
   Humidity: 96 %
   Wind Speed: 2.38 m/s
   ```

## Code Structure

- `main.go` - Main application logic to fetch and display weather data.

## API Reference

The application uses the Tomorrow.io API:

- **Endpoint:** `https://api.tomorrow.io/v4/weather/realtime`
- **Query Parameters:**
  - `location=toronto`
  - `apikey=YOUR_API_KEY`

## Troubleshooting

- If all values are `0`, check if the API key is valid and active.
- Print raw JSON response for debugging:
  ```go
  fmt.Println("Raw API Response:", string(body))
  ```

## License

This project is licensed under the MIT License.


