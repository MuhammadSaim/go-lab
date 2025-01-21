package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MuhammadSaim/go-lab/go-weather-app/config"
)

func main() {
	// laod config
	cfg := config.LoadConfig()

	// get the location in this
	var location string

	// read input line
	fmt.Print("Enter the city: ")
	fmt.Scan(&location)

	// check if location is present or not if not just terminate the system
	if location == "" {
		fmt.Println("You have to enter a location")
		return
	}

	// build the url with different params
	url := fmt.Sprintf("%s?location=%s&apikey=%s", cfg.TOMORROW_IO_API_URL, location, cfg.TOMORROW_IO_API_KEY)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Read and print raw response for debugging
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println("Weather Data for:", weather.Location.Name)
	fmt.Println("Temperature:", weather.Data.Values.Temperature, "°C")
	fmt.Println("Humidity:", weather.Data.Values.Humidity, "%")
	fmt.Println("Wind Speed:", weather.Data.Values.WindSpeed, "m/s")
}
