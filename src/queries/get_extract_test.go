package queries

import (
	"testing"
	"time"

	"gorinha/src/db"
	"gorinha/src/helpers"
	"gorinha/src/models"
)

func TestGetExtract(t *testing.T) {
	helpers.InitTestDB()
	createTestAccountAndTransactions()

	extract, err := GetExtract("1")
	if err != nil {
		t.Fatalf("Failed to get extract: %v", err)
	}

	if extract.ExtractBalance.LimitAmount != 1000.0 {
		t.Errorf("Expected LimitAmount to be 1000.0, got %v", extract.ExtractBalance.LimitAmount)
	}
	if extract.ExtractBalance.Amount != 500.0 {
		t.Errorf("Expected Amount to be 500.0, got %v", extract.ExtractBalance.Amount)
	}

	expectedTransactionCount := 2
	if len(extract.TransactionRecords) != expectedTransactionCount {
		t.Fatalf("Expected %d transactions, got %d", expectedTransactionCount, len(extract.TransactionRecords))
	}

	expectedTransactions := []TransactionRecord{
		{
			Amount:     500,
			Type:       "d",
			Description: "descricao",
			Timestamp:  time.Now(),
		},
		{
			Amount:     50,
			Type:       "c",
			Description: "descricao",
			Timestamp:  time.Now(),
		},
	}

	for i, transaction := range extract.TransactionRecords {
		if transaction.Amount != expectedTransactions[i].Amount {
			t.Errorf("Expected transaction amount to be %v, got %v", expectedTransactions[i].Amount, transaction.Amount)
		}
		if transaction.Type != expectedTransactions[i].Type {
			t.Errorf("Expected transaction type to be %v, got %v", expectedTransactions[i].Type, transaction.Type)
		}
		if transaction.Description != expectedTransactions[i].Description {
			t.Errorf("Expected transaction description to be %v, got %v", expectedTransactions[i].Description, transaction.Description)
		}
		if !transaction.Timestamp.Truncate(time.Second).Equal(expectedTransactions[i].Timestamp.Truncate(time.Second)) {
			t.Errorf("Expected transaction timestamp to be %v, got %v", expectedTransactions[i].Timestamp, transaction.Timestamp)
		}
	}

	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func createTestAccountAndTransactions() {
	account := models.Account{Name: "test account", LimitAmount: 1000}
	db.Gorm.Create(&account)
	balance := models.Balance{AccountID: account.ID, Amount: 500}
	db.Gorm.Create(&balance)

	transactions := []models.Transaction{
		{AccountID: account.ID, Amount: 50, TransactionType: "c", Description: "descricao", Date: time.Now()},
		{AccountID: account.ID, Amount: 500, TransactionType: "d", Description: "descricao", Date: time.Now()},
	}

	for _, transaction := range transactions {
		db.Gorm.Create(&transaction)
	}
}
