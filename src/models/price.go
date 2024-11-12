package models

import "time"

type Price struct {
	ID          uint      `gorm:"primaryKey"`
	Currency    string    `gorm:"unique;not null"`
	Price       float64   `gorm:"not null"`
	LastUpdated time.Time `gorm:"not null"`
}