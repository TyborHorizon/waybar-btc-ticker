package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("COINMARKETCAP_API_KEY")
	if apiKey == "" {
		log.Fatal("env COINMARKETCAP_API_KEY is missing")
	}
	coinMarketRequest(apiKey)
}

func coinMarketRequest(apiKey string) {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=BTC&convert=USD"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("request-error: %v", err)
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("invalid response %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Data map[string]struct {
			Quote map[string]struct {
				Price float64 `json:"price"`
			} `json:"quote"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("json decoding failed: %v", err)
	}

	price := result.Data["BTC"].Quote["USD"].Price
	fmt.Printf("%.2f\n", price)
}
