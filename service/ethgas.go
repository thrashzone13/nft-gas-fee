package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type EthgasService struct{}

type EthgasResponse struct {
	Fast    int `json:"fast"`
	Fastest int `json:"fastest"`
	SafeLow int `json:"safeLow"`
	Average int `json:"average"`
}

func NewEthgasService() *EthgasService {
	return &EthgasService{}
}

func (s *EthgasService) Get() *EthgasResponse {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.ethgasstation.info/json/ethgasAPI.json", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response EthgasResponse
	json.Unmarshal(data, &response)

	return &response
}
