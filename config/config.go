package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ProfileURL              string = "https://evermos.com/placeholder-profile.png"
	Config_SendbirdBaseURL  string = "SENDBIRD.BASE_URL"
	Config_SendbirdAPIToken string = "SENDBIRD.API_TOKEN"
)

type ConfigEnv struct {
	SendbirdBaseURL  string
	SendbirdAPIToken string
}

func GetSendbirdConfigVariable() (res ConfigEnv) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return ConfigEnv{
		SendbirdBaseURL:  os.Getenv(Config_SendbirdBaseURL),
		SendbirdAPIToken: os.Getenv(Config_SendbirdAPIToken),
	}

}
