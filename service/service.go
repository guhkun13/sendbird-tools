package service

import (
	"os"

	"github.com/guhkun13/sendbird-tools/config"
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
	Config         config.AppConfig
	SendirdService sendbird.Service
}

func NewService() *ServiceImpl {
	return &ServiceImpl{
		Config:         config.GetAppConfig(),
		SendirdService: sendbird.NewService(),
	}
}
