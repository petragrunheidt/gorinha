package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name        string
	LimitAmount float64
	Balances    []Balance    `gorm:"foreignKey:AccountID"`
	Transactions []Transaction `gorm:"foreignKey:AccountID"`
}
