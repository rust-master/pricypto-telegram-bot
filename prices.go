package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CoinData struct {
	ID                string   `json:"id"`
	Symbol            string   `json:"symbol"`
	Name              string   `json:"name"`
	CurrentPrice      float64  `json:"current_price"`
	MarketCap         int64    `json:"market_cap"`
	MarketCapRank     int      `json:"market_cap_rank"`
	TotalVolume       int64    `json:"total_volume"`
	High24h           float64  `json:"high_24h"`
	Low24h            float64  `json:"low_24h"`
	PriceChange24h    float64  `json:"price_change_24h"`
	PriceChangePct24h float64  `json:"price_change_percentage_24h"`
	CirculatingSupply float64  `json:"circulating_supply"`
	TotalSupply       float64  `json:"total_supply"`
	MaxSupply         *float64 `json:"max_supply"`
	ATH               float64  `json:"ath"`
	ATHChangePct      float64  `json:"ath_change_percentage"`
	ATHDate           string   `json:"ath_date"`
	ATL               float64  `json:"atl"`
	ATLChangePct      float64  `json:"atl_change_percentage"`
	ATLDate           string   `json:"atl_date"`
	ROI               struct {
		Times      float64 `json:"times"`
		Currency   string  `json:"currency"`
		Percentage float64 `json:"percentage"`
	} `json:"roi"`
	LastUpdated string `json:"last_updated"`
}

func fetchPriceData(url string) ([]CoinData, error) {

	// Make a GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into a slice of CoinData structs
	var data []CoinData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
