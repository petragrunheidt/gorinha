package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID       uint
	Amount          int `gorm:"not null"`
	TransactionType string `gorm:"not null"`
	Description     string
	Date            string
}
