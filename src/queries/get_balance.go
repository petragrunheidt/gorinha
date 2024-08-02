package queries

import (
	"fmt"
	"log"

	"gorinha/src/db"
	"gorinha/src/models"
)

type Balance struct {
	LimitAmount float64 `json:"limite"`
	Amount      float64 `json:"saldo"`
}

func GetBalance(id string) (Balance, error) {
	var account models.Account
	if err := db.DB.Preload("Balances").First(&account, id).Error; err != nil {
		return Balance{}, err
	}

	if len(account.Balances) > 0 {
		balance := account.Balances[0]
		log.Printf("added balance: %v, %v", account.LimitAmount, balance.Amount)
		return Balance{LimitAmount: account.LimitAmount, Amount: balance.Amount}, nil
	}

	return Balance{}, fmt.Errorf("no balances found for account ID %s", id)
}
