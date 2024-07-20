package commands

import (
	"fmt"

	"gorinha/src/db"
	"gorinha/src/models"

	"gorm.io/gorm"
)

func UpdateBalance(id string, amount int, transactionType string) error {
	var err error

	switch transactionType {
	case "c":
		err = db.Gorm.
		Model(&models.Account{}).
		Where("id = ?", id).
		Update("limit_amount", gorm.Expr("limit_amount - ?", amount)).
		Error
	case "d":
		err = db.Gorm.
		Model(&models.Balance{}).
		Where("account_id = ?", id).
		Update("amount", gorm.Expr("amount - ?", amount)).
		Error
	default:
		return fmt.Errorf("invalid transaction type")
	}
	return err
}
