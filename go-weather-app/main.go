package main

import (
	"fmt"
	"net/http"

	"github.com/MuhammadSaim/go-lab/go-weather-app/config"


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
	url := fmt.Sprintf("%s?location=%s&apiKey=%s", cfg.TOMORROW_IO_API_URL, location, cfg.TOMORROW_IO_API_KEY)

  // build a request the url
  req, err := http.NewRequest("GET", url, nil)

  // check if there is an error while calling the route
  if err != nil {
    fmt.Printf("Error: %v", err)
    return
  }

  // add some required headers to the request
  req.Header.Add("accept", "application/json")
	req.Header.Add("accept-encoding", "deflate, gzip, br")

  // send the request to get the data
  res, err := http.DefaultClient.Do(req)

   // check if

}
