package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold environment variables
type Config struct {
	SLACK_BOT_TOKEN  string
	SLACK_CHANNEL_ID string
}

// LoadConfig loads environment variables from the .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		SLACK_BOT_TOKEN:  os.Getenv("SLACK_BOT_TOKEN"),
		SLACK_CHANNEL_ID: os.Getenv("SLACK_CHANNEL_ID"),
	}
}
