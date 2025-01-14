package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold environment variables
type Config struct {
	SLACK_BOT_TOKEN string
	SLACK_APP_TOKEN string
}

// LoadConfig loads environment variables from the .env file
func LoadConfig() *Config {
	// load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		SLACK_BOT_TOKEN: os.Getenv("SLACK_BOT_TOKEN"),
		SLACK_APP_TOKEN: os.Getenv("SLACK_APP_TOKEN"),
	}
}
