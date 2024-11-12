package util

import (
	"log"
	"os"

	"github.com/akshaybt001/CoinMarket/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_KEY")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	db.AutoMigrate(models.Price{})
	DB = db
	log.Println("Database connected")
}
