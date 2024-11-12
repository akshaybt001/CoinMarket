package main

import (
	"log"

	"github.com/akshaybt001/CoinMarket/cronjob"
	"github.com/akshaybt001/CoinMarket/routes"
	"github.com/akshaybt001/CoinMarket/src/controllers"
	"github.com/akshaybt001/CoinMarket/src/repository"
	"github.com/akshaybt001/CoinMarket/src/service"
	"github.com/akshaybt001/CoinMarket/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	util.ConnectDatabase()
	priceRepo := repository.NewCurrencyRepository()
	priceService := &service.PriceService{Repo: priceRepo}
	priceController := &controllers.PriceController{Service: priceService}

	cronjob.StartPriceUpdaterCron(priceService)

	router := gin.Default()
	routes.SetupRoutes(router, priceController)
	router.Run(":8080")
}
