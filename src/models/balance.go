package models

import (
	"gorm.io/gorm"
)

type Balance struct {
		gorm.Model
    AccountID uint
    Amount    float64 `gorm:"not null;check:amount >= 0"`
}
