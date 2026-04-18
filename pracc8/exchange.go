package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RateResponse struct {
	Rate     float64 `json:"rate"`
	ErrorMsg string  `json:"error"`
}

type ExchangeService struct {
	BaseURL string
	Client  *http.Client
}

func NewExchangeService(url string) *ExchangeService {
	return &ExchangeService{
		BaseURL: url,
		Client:  &http.Client{},
	}
}

func (s *ExchangeService) GetRate() (float64, error) {
	resp, err := s.Client.Get(s.BaseURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var r RateResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, err
	}

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("api error")
	}

	return r.Rate, nil
}
