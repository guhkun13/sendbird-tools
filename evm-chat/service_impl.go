package evmchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (s *ServiceImpl) JoinSuperGroup(req JoinSuperGroupRequest) (res JoinSuperGroupResponse, err error) {
	funcName := "JoinSuperGroup"
	url := s.Config.Evermos.EvmChat.BaseURL + s.Endpoints.JoinSuperGroup
	fmt.Printf(">> [%s] %s to = %s", funcName, http.MethodPost, url)

	jsonData, err := json.Marshal(req)
	payload := bytes.NewBuffer(jsonData)

	request, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		panic(err)
	}

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

	log.Info().Interface("res", res).Msg("[JoinSuperGroupResponse]")

	return
}
