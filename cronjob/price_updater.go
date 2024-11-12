package cronjob

import (
	"log"

	"github.com/akshaybt001/CoinMarket/src/service"
	"github.com/robfig/cron"
)

func StartPriceUpdaterCron(priceService *service.PriceService) {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		if err := priceService.UpdatePrices(); err != nil {
			log.Println("Failed to update prices:", err)
		}
	})
	c.Start()
}
