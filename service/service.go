package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/guhkun13/sendbird-tools/config"
	"github.com/guhkun13/sendbird-tools/sendbird"
	"github.com/rs/zerolog/log"
)

var (
	FuncOnboardingUser     string = "OnboardingUser"
	FuncCreateGroupChannel string = "CreateGroupChannel"
	FuncFreezeChannel      string = "FreezeChannel"
	FuncSendWelcomeMessage string = "SendWelcomeMessage"
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

// onboarding:
// [1] create channel,
// [2] freeze channel,
// [3] send welcome message
func OnboardingUser(req WorkerRequest) {
	funcName := "OnboardingUser"
	log.Info().Str("func", funcName).Msg("[Main Flow]")

	data := req.Users
	for idx, user := range data {
		fmt.Printf("[ idx: %d | userID: %s ] \n", idx, user.UserID)

		// [1] Create Group Channel
		reqCreateChannel, resCreateChannel := CreateGroupChannel(user.UserID)

		// save log
		logCreate := HttpLog{
			Index:    idx,
			Function: FuncCreateGroupChannel,
			Request:  reqCreateChannel,
			Response: resCreateChannel,
		}
		WriteLog(logCreate, req.LogFile)

		// [2] Freeze Group Channel
		reqFreezeChannel, resFreezeChannel := FreezeGroupChannel(user.UserID)
		// save log
		logFreeze := HttpLog{
			Index:    idx,
			Function: FuncFreezeChannel,
			Request:  reqFreezeChannel,
			Response: resFreezeChannel,
		}
		WriteLog(logFreeze, req.LogFile)

		// [3] Send Welcome Messsage
		if resCreateChannel.Code == http.StatusOK {
			reqSendMessage, resSendMessage := SendWelcomeMessage(user.UserID)

			// save log
			logSend := HttpLog{
				Index:    idx,
				Function: FuncSendWelcomeMessage,
				Request:  reqSendMessage,
				Response: resSendMessage,
			}
			WriteLog(logSend, req.LogFile)
			checkDelay()
		}
		// check for delay
		checkDelay()
	}
}

// ONE
func CreateGroupChannel(userID string) (req interface{}, res sendbird.HttpResponse) {
	// funcName := "CreateGroupChannel"
	// log.Info().
	// 	Str("func", funcName).
	// 	Msg("[Main Flow]")

	userIDint, _ := strconv.ParseInt(userID, 10, 64)
	reqSendbird := sendbird.CreateGroupChannelRequest{
		Name:       "Evermos Info",
		ChannelURL: "evm_info_" + userID,
		CoverURL:   SendbirdGroupCoverlURL,
		CustomType: "evm_info",
		IsDistinct: false,
		UserIDs:    []int64{userIDint},
		Data:       "Sumber informasi terkini dan menarik di Evermos khusus untukmu.",
	}

	// send request
	req = reqSendbird
	res, err := sendbird.CreateGroupChannel(reqSendbird)
	if err != nil {
		fmt.Printf("err: %v \n", err)
	}

	return
}

func FreezeGroupChannel(userID string) (req interface{}, res sendbird.HttpResponse) {
	reqSendbird := sendbird.FreezeGroupChannelRequest{
		ChannelURL: "evm_info_" + userID,
		Freeze:     true,
	}

	// send request
	req = reqSendbird
	res, err := sendbird.FreezeGroupChannel(reqSendbird)
	if err != nil {
		fmt.Printf("err: %v \n", err)
	}

	return
}

func SendWelcomeMessage(userID string) (req interface{}, res sendbird.HttpResponse) {
	reqSendbird := sendbird.SendMessageRequest{
		MessageType:         "ADMM",
		ChannelURL:          "evm_info_" + userID,
		CustomType:          "evm_info",
		Message:             EvmInfoWelcomeMessage,
		PushMessageTemplate: sendbird.PushMessageTemplate,
		SendPush:            fmt.Sprintf("%v", sendbird.SendPush),
		MarkAsRead:          fmt.Sprintf("%v", sendbird.MarkAsRead),
		IsSilent:            fmt.Sprintf("%v", sendbird.IsSilent),
	}

	req = reqSendbird
	res, err := sendbird.SendMessage(reqSendbird)
	if err != nil {
		fmt.Printf("err: %v \n", err)
	}
	return
}

// BULK
func BulkCreateGroupChannel(req WorkerRequest) {
	funcName := "BulkCreateGroupChannel"
	log.Info().
		Str("func", funcName).
		Msg("[Main Flow]")

	data := req.Users
	for idx, user := range data {
		fmt.Printf("idx: %d", idx)
		reqExt, resExt := CreateGroupChannel(user.UserID)

		// save log
		dataLog := HttpLog{
			Index:    idx,
			Function: funcName,
			Request:  reqExt,
			Response: resExt,
		}
		WriteLog(dataLog, req.LogFile)
		checkDelay()
	}
}

func BulkSendWelcomeMessage(req WorkerRequest) {
	funcName := "BulkSendWelcomeMessage"
	data := req.Users
	for idx, user := range data {
		fmt.Print("idx: ", idx)
		reqExt, resExt := SendWelcomeMessage(user.UserID)

		// save log
		dataLog := HttpLog{
			Index:    idx,
			Function: funcName,
			Request:  reqExt,
			Response: resExt,
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
