package sendbird

import (
	"github.com/guhkun13/sendbird-tools/config"
)

type Service interface {
	CreateGroupChannel(req CreateGroupChannelRequest) (res HttpResponse, err error)
	FreezeGroupChannel(req FreezeGroupChannelRequest) (res HttpResponse, err error)
	SendMessage(req SendMessageRequest) (res HttpResponse, err error)
}
type ServiceImpl struct {
	Config    config.Config
	Endpoints Endpoints
}
type Endpoints struct {
	SendMessage        string
	CreateGroupChannel string
	FreezeGroupChannel string
}

func InitService() *ServiceImpl {
	return &ServiceImpl{
		Config: config.InitConfig(),
		Endpoints: Endpoints{
			SendMessage:        "/v3/group_channels/{channel_url}/messages",
			CreateGroupChannel: "/v3/group_channels",
			FreezeGroupChannel: "/v3/group_channels/{channel_url}/freeze",
		},
	}
}
