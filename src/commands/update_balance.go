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

func UpdateBalance(id string, amount float64, transactionType string, description string) error {
	tx := db.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	var err error
	err = updateTransaction(tx, id, amount, transactionType)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction failed: %w", err)
	}

	err = registerTransaction(tx, id, amount, transactionType, description)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func updateTransaction(tx *gorm.DB, id string, amount float64, transactionType string) error {
	var err error

	switch transactionType {
	case "c":
		var account models.Account
		if err := tx.First(&account, id).Error; err != nil {
			return err
		}
		account.LimitAmount -= amount
		err = tx.Save(&account).Error
	case "d":
		var balance models.Balance
		if err := tx.Where("account_id = ?", id).First(&balance).Error; err != nil {
			return err
		}
		balance.Amount -= amount
		err = tx.Save(&balance).Error
	default:
		return fmt.Errorf("invalid transaction type")
	}

	return err
}

func registerTransaction(tx *gorm.DB, id string, amount float64, transactionType string, description string) error {
	newTransaction := models.Transaction{
		AccountID:       parseUint(id),
		Amount:          amount,
		TransactionType: transactionType,
		Description:     description,
		Date:            time.Now(),
	}

	result := tx.Create(&newTransaction)
	if result.Error != nil {
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
