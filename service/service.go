package service

import (
	"os"

	"github.com/guhkun13/sendbird-tools/config"
	evmchat "github.com/guhkun13/sendbird-tools/evm-chat"
	"github.com/guhkun13/sendbird-tools/sendbird"
)

type Service interface {
	CreateUserList(data [][]string) (res MigratedUserSendbirdList)
	OnboardingUser(req WorkerRequest)
	CreateGroupChannel(userID string) (req interface{}, res sendbird.HttpResponse)
	FreezeGroupChannel(userID string) (req interface{}, res sendbird.HttpResponse)
	SendWelcomeMessage(userID string) (req interface{}, res sendbird.HttpResponse)
	// log
	CreateLogFile(csvFile string) (res *os.File)
	WriteLog(data HttpLog, f *os.File)
}

type ServiceImpl struct {
	Config             config.Config
	SendirdService     sendbird.Service
	EvermosChatService evmchat.Service
}

func InitService() *ServiceImpl {
	return &ServiceImpl{
		Config:             config.InitConfig(),
		SendirdService:     sendbird.InitService(),
		EvermosChatService: evmchat.InitService(),
	}
}
