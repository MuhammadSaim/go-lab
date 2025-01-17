package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold environment variables
type Config struct {
	TOMORROW_IO_API_KEY string
	TOMORROW_IO_API_URL string
}

// LoadConfig loads environment variables from the .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		TOMORROW_IO_API_KEY: os.Getenv("TOMORROW_IO_API_KEY"),
		TOMORROW_IO_API_URL: "https://api.tomorrow.io/v4/weather/realtime",
	}
}
