package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://www.ethgasstation.info/json/ethgasAPI.json", nil)
	if err != nil {
		log.Panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var response EthgasResponse
	json.Unmarshal(data, &response)

	return &response
}
