package models

import (
	"time"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID       uint
	Amount          float64 `gorm:"not null"`
	TransactionType string `gorm:"not null"`
	Description     string
	Date            time.Time
}
