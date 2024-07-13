package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID       uint
	Amount          int
	TransactionType string
	Description     string
	Date            string
}
