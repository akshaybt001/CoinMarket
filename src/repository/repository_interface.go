package repository

import (
	"time"

	"github.com/akshaybt001/CoinMarket/src/models"
)

type CurrencyRespository interface {
	UpdateCurrencyPrice(currency string, price float64, lastUpdated time.Time) error
	GetAllPrices() ([]models.Price, error)
}