package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CryptonatorService struct{}

type CryptonatorResponse struct {
	ETH float64 `json:"ETH"`
}

func NewCryptonatorService() *CryptonatorService {
	return &CryptonatorService{}
}

func (s *CryptonatorService) Get() *CryptonatorResponse {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://min-api.cryptocompare.com/data/price?fsym=USD&tsyms=ETH", nil)
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

	var response CryptonatorResponse
	json.Unmarshal(data, &response)

	return &response
}
