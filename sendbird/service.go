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
	HTTP_GET  string = "GET"
	HTTP_POST string = "POST"
	HTTP_PUT  string = "PUT"
)

type Service interface {
	CreateGroupChannel(req CreateGroupChannelRequest) (res HttpResponse, err error)
	FreezeGroupChannel(req FreezeGroupChannelRequest) (res HttpResponse, err error)
	SendMessage(req SendMessageRequest) (res HttpResponse, err error)
}

type ServiceImpl struct {
	Config    config.SendbirdConfig
	Endpoints Endpoints
}
type Endpoints struct {
	SendMessage        string
	CreateGroupChannel string
	FreezeGroupChannel string
}

func NewService() *ServiceImpl {
	return &ServiceImpl{
		Config: config.GetSendbirdConfig(),
		Endpoints: Endpoints{
			SendMessage:        "/v3/group_channels/{channel_url}/messages",
			CreateGroupChannel: "/v3/group_channels",
			FreezeGroupChannel: "/v3/group_channels/{channel_url}/freeze",
		},
	}
}

func (s *ServiceImpl) CreateGroupChannel(req CreateGroupChannelRequest) (res HttpResponse, err error) {
	funcName := "CreateGroupChannel"
	fmt.Printf(">> [%s] %s to = %s", funcName, HTTP_POST, req.ChannelURL)

	url := s.Config.SendbirdBaseURL + s.Endpoints.CreateGroupChannel
	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_POST, url, payload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Api-Token", s.Config.SendbirdAPIToken)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("|| status [%d]: %s \n", response.StatusCode, response.Status)
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Cannont unmarshal response")
	}
	res.Status = response.Status
	res.Code = response.StatusCode
	res.Body = string(body)

	return
}

func (s *ServiceImpl) FreezeGroupChannel(req FreezeGroupChannelRequest) (res HttpResponse, err error) {
	funcName := "FreezeGroupChannel"
	fmt.Printf(">> [%s] %s  to = %s", funcName, HTTP_PUT, req.ChannelURL)

	url := s.Config.SendbirdBaseURL + s.Endpoints.FreezeGroupChannel
	url = strings.Replace(url, "{channel_url}", req.ChannelURL, -1)
	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_PUT, url, payload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Api-Token", s.Config.SendbirdAPIToken)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("|| status [%d]: %s \n", response.StatusCode, response.Status)

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Cannont unmarshal response")
	}
	res.Status = response.Status
	res.Code = response.StatusCode
	res.Body = string(body)

	return
}

func (s *ServiceImpl) SendMessage(req SendMessageRequest) (res HttpResponse, err error) {
	funcName := "SendMessage"
	fmt.Printf(">> [%s] %s to = %s", funcName, HTTP_POST, req.ChannelURL)

	url := s.Config.SendbirdBaseURL + s.Endpoints.SendMessage
	url = strings.Replace(url, "{channel_url}", req.ChannelURL, -1)
	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_POST, url, payload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Api-Token", s.Config.SendbirdAPIToken)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Printf("|| status [%d]: %s \n", response.StatusCode, response.Status)

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Cannont unmarshal response")
	}
	res.Status = response.Status
	res.Code = response.StatusCode
	res.Body = string(body)

	return
}
