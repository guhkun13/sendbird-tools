package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/guhkun13/sendbird-tools/config"
	"github.com/guhkun13/sendbird-tools/sendbird"
	"github.com/rs/zerolog/log"
)

func CreateUserList(data [][]string) (res MigratedUserSendbirdList) {
	for i, line := range data {
		if i > 0 { // omit header line
			var rec MigratedUserSendbird
			for j, field := range line {
				if j == 0 {
					rec.UserID = field
				}
			}
			res = append(res, rec)
		}
	}

	return
}

func CreateGroupChannel(req WorkerRequest) {
	funcName := "CreateGroupChannel"
	log.Info().
		Str("func", "CreateGroupChannel").
		Msg("[Main Flow]")

	data := req.Users
	for idx, user := range data {
		fmt.Print("idx: ", idx)
		userIDint, _ := strconv.ParseInt(user.UserID, 10, 64)

		reqSendbird := sendbird.CreateGroupChannelRequest{
			Name:       "Evermos Info",
			ChannelURL: "evm_info_" + user.UserID,
			CustomType: "evm_info",
			IsDistinct: false,
			UserIDs:    []int64{userIDint},
			Data:       "Sumber informasi terkini dan menarik di Evermos khusus untukmu.",
		}

		// send request
		res, err := sendbird.CreateGroupChannel(reqSendbird)
		if err != nil {
			fmt.Printf("err: %v \n", err)
		}

		// save log
		dataLog := SendLog{
			Index:    idx,
			Function: funcName,
			Request:  reqSendbird,
			Response: res,
		}
		WriteLog(dataLog, req.LogFile)
		checkDelay()
	}
}

func SendWelcomeMessage(req WorkerRequest) {
	funcName := "SendWelcomeMessage"
	data := req.Users
	for idx, user := range data {
		fmt.Print("idx: ", idx)
		// prepare request
		reqSendbird := sendbird.SendMessageRequest{
			MessageType:         "ADMM",
			ChannelURL:          "evm_info_" + user.UserID,
			CustomType:          "evm_info",
			Message:             EvmInfoWelcomeMessage,
			PushMessageTemplate: sendbird.PushMessageTemplate,
			SendPush:            fmt.Sprintf("%v", sendbird.SendPush),
			MarkAsRead:          fmt.Sprintf("%v", sendbird.MarkAsRead),
			IsSilent:            fmt.Sprintf("%v", sendbird.IsSilent),
		}

		// send request
		res, err := sendbird.SendMessage(reqSendbird)
		if err != nil {
			fmt.Printf("err: %v \n", err)
		}

		// save log
		dataLog := SendLog{
			Index:    idx,
			Function: funcName,
			Request:  reqSendbird,
			Response: res,
		}
		WriteLog(dataLog, req.LogFile)
		checkDelay()
	}
}

func checkDelay() {
	isDelay, _ := strconv.ParseBool(config.GetAppConfig().AppDelayEnabled)
	delayTime, _ := strconv.ParseInt(config.GetAppConfig().AppDelayTime, 10, 64)

	if isDelay {
		log.Info().Int64("time", delayTime).Msg("[Delay]")
		time.Sleep(time.Millisecond * time.Duration(delayTime))
	}
}
