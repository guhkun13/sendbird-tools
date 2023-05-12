package sendbird

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/guhkun13/sendbird-tools/config"
)

const (
	HTTP_GET  = "GET"
	HTTP_POST = "POST"
)

type Endpoints struct {
	SendMessage string
	CreateGroup string
}

func GetEndpoints() Endpoints {
	return Endpoints{
		SendMessage: "/v3/group_channels/{channel_url}/messages",
		CreateGroup: "/v3/group_channels",
	}
}

func CreateGroupChannel(req CreateGroupChannelRequest) (res CreateGroupChannelResponse, err error) {
	conf := config.GetSendbirdConfig()

	postUrl := conf.SendbirdBaseURL + GetEndpoints().CreateGroup
	fmt.Printf("||| %s to = %s", HTTP_POST, req.ChannelURL)

	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_POST, postUrl, payload)
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

	fmt.Println("||| status:", response.Status)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Cannont unmarshal response")
	}

	return
}

func SendMessage(req SendMessageRequest) (res SendMessageResponse, err error) {
	conf := config.GetSendbirdConfig()

	postUrl := conf.SendbirdBaseURL + GetEndpoints().SendMessage
	postUrl = strings.Replace(postUrl, "{channel_url}", req.ChannelURL, -1)
	fmt.Print("||| Send to = ", req.ChannelURL)

	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_POST, postUrl, payload)
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

	fmt.Println("||| status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Cannont unmarshal response")
	}
	res.Status = response.Status

	return
}
