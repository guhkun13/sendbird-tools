// source: https://gosamples.dev/read-csv/
package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// DTO
const EvmInfoWelcomeMessage = `Assalamu'alaikum Reseller,

Selamat datang di channel Evermos Info!

Kini, kamu makin mudah dapatkan info menarik terkait Evermos. Info yang akan didapatkan tentunya beragam, dan membantu perjuangan ikhtiarmu. Jadi, sering-sering ya intip info channel ini supaya tidak ketinggalan kabar terkini.

Salam Sungkem dari Kami,
Seluruh Tim Evermos`

const (
	ProfileURL              string = "https://evermos.com/placeholder-profile.png"
	Config_SendbirdBaseURL  string = "SENDBIRD.BASE_URL"
	Config_SendbirdAPIToken string = "SENDBIRD.API_TOKEN"
)

var SendbirdEndpointSendMessage string = "/v3/group_channels/{channel_url}/messages"

type ConfigEnv struct {
	SendbirdBaseURL  string
	SendbirdAPIToken string
}
type MigratedUserSendbird struct {
	UserID   string
	FullName string
}

type MigratedUserSendbirdList []MigratedUserSendbird

type SendLog struct {
	Index    int
	Request  interface{}
	Response interface{}
}

type BlastWelcomeMessageRequest struct {
	Users   MigratedUserSendbirdList
	LogFile *os.File
}

// sendbird
const (
	PushMessageTemplate string = "alternative"
	SendPush            bool   = true
	IsSilent            bool   = false
	MarkAsRead          bool   = false
)

type SendMessageRequest struct {
	MessageType         string `json:"message_type"`
	ChannelURL          string `json:"channel_url"`
	CustomType          string `json:"custom_type"`
	UserId              string `json:"user_id"`
	Message             string `json:"message"`
	Data                string `json:"data,omitempty"`
	SendPush            string `json:"send_push,omitempty"`
	PushMessageTemplate string `json:"push_message_template,omitempty"`
	IsSilent            string `json:"is_silent,omitempty"`
	MarkAsRead          string `json:"mark_as_read,omitempty"`
}

type SendMessageResponse struct {
	Status     string `json:"status"`
	Type       string `json:"type"`
	ChannelURL string `json:"channel_url"`
	CustomType string `json:"custom_type"`
	MessageID  int64  `json:"message_id"`
	Message    string `json:"message"`
	CreatedAt  int64  `json:"created_at"`
}

// main
func main() {
	// open file
	csvFile := os.Args[1]

	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("error: %v \n", err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	userList := createUserList(data)
	logFile := createLogFile(csvFile)
	req := BlastWelcomeMessageRequest{
		Users:   userList,
		LogFile: logFile,
	}
	blastWelcomeMessage(req)
	defer logFile.Close()
}

// blasting
func blastWelcomeMessage(req BlastWelcomeMessageRequest) {
	data := req.Users[:3]
	for idx, user := range data {
		// prepare request
		reqSendbird := SendMessageRequest{
			MessageType:         "ADMM",
			ChannelURL:          "evm_info_" + user.UserID,
			CustomType:          "evm_info",
			Message:             EvmInfoWelcomeMessage,
			PushMessageTemplate: PushMessageTemplate,
			SendPush:            fmt.Sprintf("%v", SendPush),
			MarkAsRead:          fmt.Sprintf("%v", MarkAsRead),
			IsSilent:            fmt.Sprintf("%v", IsSilent),
		}
		// send request
		res, err := sendMessage(reqSendbird)
		if err != nil {
			fmt.Printf("err: %v \n", err)
		}

		// save log
		dataLog := SendLog{
			Index:    idx,
			Request:  user,
			Response: res,
		}
		writeLog(dataLog, req.LogFile)
	}
}

func sendMessage(req SendMessageRequest) (res SendMessageResponse, err error) {
	conf := getConfigVariable()
	fmt.Println("conf = ", conf)

	postUrl := conf.SendbirdBaseURL + SendbirdEndpointSendMessage
	postUrl = strings.Replace(postUrl, "{channel_url}", req.ChannelURL, -1)
	fmt.Println("||| Sending to = ", postUrl)

	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest("POST", postUrl, payload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Api-Token", conf.SendbirdAPIToken)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Cannont unmarshal response")
	}
	res.Status = response.Status

	return
}

// logging
func createLogFile(csvFile string) (res *os.File) {
	logDir := "log"
	// ignore the error
	_ = os.Mkdir(logDir, os.ModePerm)

	ts := time.Now().Format("2006-01-02-15-04")
	strTs := fmt.Sprintf("%s", ts)
	outFile := logDir + "/input_" + csvFile + "_" + strTs + ".log"

	res, err := os.OpenFile(outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("output file = %v \n", outFile)
	return
}

func writeLog(data SendLog, f *os.File) {
	dataByte, _ := json.MarshalIndent(data, "", " ")

	n, err := f.Write(dataByte)
	if err != nil {
		fmt.Println(n, err)
	}
	if n, err = f.WriteString(",\n"); err != nil {
		fmt.Println(n, err)
	}
}

// Users
func createUserList(data [][]string) (res MigratedUserSendbirdList) {
	for i, line := range data {
		if i > 0 { // omit header line
			var rec MigratedUserSendbird
			for j, field := range line {
				if j == 0 {
					rec.UserID = field
				} else if j == 1 {
					rec.FullName = field
				}
			}
			res = append(res, rec)
		}
	}

	return
}

func getConfigVariable() (res ConfigEnv) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return ConfigEnv{
		SendbirdBaseURL:  os.Getenv(Config_SendbirdBaseURL),
		SendbirdAPIToken: os.Getenv(Config_SendbirdAPIToken),
	}

}
