package controllers

import (
	"net/http"

	"github.com/akshaybt001/CoinMarket/src/service"
	"github.com/gin-gonic/gin"
)

type PriceController struct {
	Service *service.PriceService
}

func (ctrl *PriceController) GetAllPrices(c *gin.Context) {
	prices, err := ctrl.Service.Repo.GetAllPrices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve prices"})
		return
	}
	c.JSON(http.StatusOK, prices)
}