package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	AppDelayEnabled string = "APP.DELAY.ENABLED"
	AppDelayTime    string = "APP.DELAY.TIME"
)

const (
	SendbirdProfileURL string = "https://evermos.com/placeholder-profile.png"
	SendbirdBaseURL    string = "SENDBIRD.BASE_URL"
	SendbirdAPIToken   string = "SENDBIRD.API_TOKEN"
)

type AppConfig struct {
	AppDelayEnabled string
	AppDelayTime    string
}

type SendbirdConfig struct {
	SendbirdBaseURL  string
	SendbirdAPIToken string
}

// app config
func GetAppConfig() (res AppConfig) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return AppConfig{
		AppDelayEnabled: os.Getenv(AppDelayEnabled),
		AppDelayTime:    os.Getenv(AppDelayTime),
	}

}

// sendbirf config
func GetSendbirdConfig() (res SendbirdConfig) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return SendbirdConfig{
		SendbirdBaseURL:  os.Getenv(SendbirdBaseURL),
		SendbirdAPIToken: os.Getenv(SendbirdAPIToken),
	}

}
