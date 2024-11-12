package repository

import (
	"time"

	"github.com/akshaybt001/CoinMarket/src/models"
	"github.com/akshaybt001/CoinMarket/util"
)

type currencyRepo struct{}

func NewCurrencyRepository() CurrencyRespository {
	return &currencyRepo{}
}

func (repo *currencyRepo) UpdateCurrencyPrice(currency string, price float64, lastUpdated time.Time) error {
	return util.DB.Exec("INSERT INTO prices (currency,price,last_updated) VALUES(?,?,?) ON CONFLICT (currency) DO UPDATE SET price=EXCLUDED.price,last_updated = EXCLUDED.last_updated",
		currency, price, lastUpdated).Error
}

func (repo *currencyRepo) GetAllPrices() ([]models.Price, error) {
	var prices []models.Price
	if err := util.DB.Find(&prices).Error; err != nil {
		return nil, err
	}
	return prices, nil
}
