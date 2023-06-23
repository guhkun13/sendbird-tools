package evmchat

import (
	"errors"

	"github.com/guhkun13/sendbird-tools/config"
)

var (
	ErrFromEvmChat = errors.New("error from evm-chat")
)

type Service interface {
	JoinSuperGroup(req JoinSuperGroupRequest) (res JoinSuperGroupResponse, err error)
}

type Endpoints struct {
	JoinSuperGroup string
}

type ServiceImpl struct {
	Config    config.Config
	Endpoints Endpoints
}

func InitService() *ServiceImpl {
	return &ServiceImpl{
		Config: config.InitConfig(),
		Endpoints: Endpoints{
			JoinSuperGroup: "/v1/internal/channel/join-super-group",
		},
	}
}
