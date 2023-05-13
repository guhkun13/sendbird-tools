package sendbird

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	HTTP_GET  string = "GET"
	HTTP_POST string = "POST"
	HTTP_PUT  string = "PUT"
)

func (s *ServiceImpl) CreateGroupChannel(req CreateGroupChannelRequest) (res HttpResponse, err error) {
	funcName := "CreateGroupChannel"
	fmt.Printf(">> [%s] %s to = %s", funcName, HTTP_POST, req.ChannelURL)

	url := s.Config.Sendbird.BaseURL + s.Endpoints.CreateGroupChannel
	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_POST, url, payload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Api-Token", s.Config.Sendbird.APIToken)

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

	url := s.Config.Sendbird.BaseURL + s.Endpoints.FreezeGroupChannel
	url = strings.Replace(url, "{channel_url}", req.ChannelURL, -1)
	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_PUT, url, payload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Api-Token", s.Config.Sendbird.APIToken)

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

	url := s.Config.Sendbird.BaseURL + s.Endpoints.SendMessage
	url = strings.Replace(url, "{channel_url}", req.ChannelURL, -1)
	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(HTTP_POST, url, payload)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Api-Token", s.Config.Sendbird.APIToken)

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
