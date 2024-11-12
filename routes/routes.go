package routes

import (
	"github.com/akshaybt001/CoinMarket/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, priceController *controllers.PriceController) {
	router.GET("/getPrice",priceController.GetAllPrices)
}