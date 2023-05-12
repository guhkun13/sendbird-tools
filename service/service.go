package service

import (
	"fmt"

	"github.com/guhkun13/sendbird-tools/sendbird"
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

func SendWelcomeMessage(req WorkerRequest) {

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
			Request:  user,
			Response: res,
		}
		WriteLog(dataLog, req.LogFile)
		// time.Sleep(time.Millisecond * 200)
	}
}
