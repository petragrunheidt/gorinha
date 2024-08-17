package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID       uint
	Amount          int    `gorm:"not null"`
	TransactionType string `gorm:"not null"`
	Description     string `gorm:"not null"`
	Date            time.Time
}

func (t *Transaction) BeforeSave(tx *gorm.DB) (err error) {
	if len(t.Description) > 10 || len(t.Description) < 1 {
		return errors.New("description must have between 1 and 10 characters")
	}

	if t.TransactionType != "c" && t.TransactionType != "d" {
		return errors.New("transaction type must be either 'c' or 'd'")
	}

	return nil
}
