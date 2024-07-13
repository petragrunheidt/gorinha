package queries

import (
	"fmt"
	"log"

	"gorinha/src/db"
	"gorinha/src/models"
)

type Balance struct {
	LimitAmount float64 `json:"limit_amount"`
	Amount      float64 `json:"amount"`
}

func GetBalance(id uint) (Balance, error) {
	var account models.Account
	if err := db.Gorm.Preload("Balances").First(&account, id).Error; err != nil {
			return Balance{}, err
	}

	if len(account.Balances) > 0 {
			balance := account.Balances[0]
			log.Printf("added balance: %v, %v", account.LimitAmount, balance.Amount)
			return Balance{LimitAmount: account.LimitAmount, Amount: balance.Amount}, nil
	}

	return Balance{}, fmt.Errorf("No balances found for account ID %d", id)
}
