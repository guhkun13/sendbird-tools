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

func SendMessage(req SendMessageRequest) (res SendMessageResponse, err error) {
	conf := config.GetSendbirdConfigVariable()

	postUrl := conf.SendbirdBaseURL + EndpointSendMessage
	postUrl = strings.Replace(postUrl, "{channel_url}", req.ChannelURL, -1)
	fmt.Print("||| Send to = ", req.ChannelURL)

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