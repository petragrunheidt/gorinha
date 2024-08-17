package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name        string `gorm:"not null"`
	LimitAmount int `gorm:"not null;check:limit_amount >= 0"`
	Balances    []Balance    `gorm:"foreignKey:AccountID"`
	Transactions []Transaction `gorm:"foreignKey:AccountID"`
}
