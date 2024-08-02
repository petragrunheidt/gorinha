package commands

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gorinha/src/db"
	"gorinha/src/models"

	"gorm.io/gorm"
)

func UpdateBalance(id string, amount float64, transactionType string) error {
	var err error
	tx := db.DB.Begin()

	defer func() {
		if r:= recover(); r!= nil {
			tx.Rollback()
			panic(r)
		}
	}()

	err = updateTransaction(id, amount, transactionType)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction failed: %w", err)
	}

	err = registerTransaction(id, amount, transactionType)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return err
}

func updateTransaction(id string, amount float64, transactionType string) error {
	var err error

	switch transactionType {
	case "c":
		err = db.DB.
			Model(&models.Account{}).
			Where("id = ?", id).
			Update("limit_amount", gorm.Expr("limit_amount - ?", amount)).
			Error
	case "d":
		err = db.DB.
			Model(&models.Balance{}).
			Where("account_id = ?", id).
			Update("amount", gorm.Expr("amount - ?", amount)).
			Error
	default:
		return fmt.Errorf("invalid transaction type")
	}

	if err != nil {
		return err
	}
	return nil
}

func registerTransaction(id string, amount float64, transactionType string) error {
	newTransaction := models.Transaction{
		AccountID:       parseUint(id),
		Amount:          amount,
		TransactionType: transactionType,
		Description:     "New Transaction",
		Date:            time.Now(),
	}

	result := db.DB.Create(&newTransaction)
	if result.Error != nil {
		log.Fatal("failed to create transaction", result.Error)
		return result.Error
	}

	fmt.Println("New transaction created with ID:", newTransaction.ID)
	return nil
}

func parseUint(str string) uint {
	uintValue, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Fatalf("Failed to parse uint: %v", err)
	}
	return uint(uintValue)
}
