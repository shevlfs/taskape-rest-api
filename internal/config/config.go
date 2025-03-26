package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	BackendHost string

	TwilioAccountSID  string
	TwilioAuthToken   string
	TwilioPhoneNumber string

	Environment string
	Debug       bool
}

func Load() (*Config, error) {
	godotenv.Load()

	debug, _ := strconv.ParseBool(getEnv("DEBUG", "false"))

	return &Config{
		Port: getEnv("PORT", "8080"),

		BackendHost: getEnv("BACKEND_HOST", "localhost:50051"),

		TwilioAccountSID:  getEnv("TWILIO_ACCOUNT_SID", ""),
		TwilioAuthToken:   getEnv("TWILIO_AUTH_TOKEN", ""),
		TwilioPhoneNumber: getEnv("TWILIO_PHONE_NUMBER", ""),

		Environment: getEnv("ENV", "development"),
		Debug:       debug,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
