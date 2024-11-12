package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akshaybt001/CoinMarket/src/repository"
)

type PriceServiceInterface interface {
	UpdatePrices() error
}

type PriceService struct {
	Repo repository.CurrencyRespository
}

func (s *PriceService) UpdatePrices() error {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
	req, _ := http.NewRequest("GET", url, nil) 
	apiKey := os.Getenv("APIKEY")
	req.Header.Set("X-CMC_PRO_API_KEY", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		// Read the response body and print it for diagnostics
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to fetch data: status %d, body: %s", resp.StatusCode, string(body))
	}

	var response struct {
		Data []struct {
			Symbol string `json:"symbol"`
			Quote  struct {
				USD struct {
					Price float64 `json:"price"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"data"`
	}

	// Decode JSON response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	fmt.Println(response) // For debugging purposes
	for _, currency := range response.Data {
		if err := s.Repo.UpdateCurrencyPrice(currency.Symbol, currency.Quote.USD.Price, time.Now()); err != nil {
			log.Println("Failed to upsert currency:", currency.Symbol, err)
		}
	}
	return nil
}
