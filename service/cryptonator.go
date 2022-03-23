package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CryptonatorService struct{}

type CryptonatorResponse struct {
	ETH float64 `json:"ETH"`
}

func NewCryptonatorService() *CryptonatorService {
	return &CryptonatorService{}
}

func (s *CryptonatorService) Get() *CryptonatorResponse {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://min-api.cryptocompare.com/data/price?fsym=USD&tsyms=ETH", nil)
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

	var response CryptonatorResponse
	json.Unmarshal(data, &response)

	return &response
}
