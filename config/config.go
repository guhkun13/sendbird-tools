package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Delay struct {
			Enabled string
			Time    string
		}
	}
	Sendbird struct {
		BaseURL  string
		APIToken string
	}
	Evermos struct {
		EvmChat struct {
			BaseURL string
		}
	}
}

var (
	// app
	AppDelayEnabled string = "APP.DELAY.ENABLED"
	AppDelayTime    string = "APP.DELAY.TIME"
	// sendbird
	SendbirdBaseURL  string = "SENDBIRD.BASE_URL"
	SendbirdAPIToken string = "SENDBIRD.API_TOKEN"
	// evermos
	EvermosEvmChatBaseURL string = "EVERMOS.EVM_CHAT.BASE_URL"
)

func InitConfig() (conf Config) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// app
	conf.App.Delay.Enabled = os.Getenv(AppDelayEnabled)
	conf.App.Delay.Time = os.Getenv(AppDelayTime)
	// sendbird
	conf.Sendbird.BaseURL = os.Getenv(SendbirdBaseURL)
	conf.Sendbird.APIToken = os.Getenv(SendbirdAPIToken)
	// evermos
	conf.Evermos.EvmChat.BaseURL = os.Getenv(SendbirdBaseURL)

	return
}
