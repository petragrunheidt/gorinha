package models

import (
	"errors"
	"math"

	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	AccountID uint
	Amount    float64 `gorm:"not null;"`
}

func (b *Balance) BeforeSave(tx *gorm.DB) (err error) {
	var account Account
	if err := tx.First(&account, "id = ?", b.AccountID).Error; err != nil {
		return err
	}

	if b.Amount < 0 && math.Abs(b.Amount) > account.LimitAmount {
		return errors.New("balance amount exceeds the account's limit amount")
	}

	return nil
}
