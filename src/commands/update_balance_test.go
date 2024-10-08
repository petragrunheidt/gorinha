package commands

import (
	"testing"
	"time"

	"gorinha/src/db"
	"gorinha/src/helpers"
	"gorinha/src/models"
	"gorinha/src/queries"
)

func TestUpdateBalance(t *testing.T) {
	defer db.Close()
	t.Run("Credit", func(t *testing.T) {
		prepareDb()

		// assert initial balance
		assertBalance(t, 1000.0, 500.0)

		err := UpdateBalance("1", 500, "c", "descript")
		if err != nil {
			t.Fatalf("Failed to update balance: %v", err)
		}

		// assert updated balance
		assertBalance(t, 500.0, 500.0)
		// assert that the transaction was created
		assertTransactionCreated(t, "1", 500, "c", "descript")
	})
	t.Run("Debit", func(t *testing.T) {
		prepareDb()

		// assert initial balance
		assertBalance(t, 1000.0, 500.0)

		err := UpdateBalance("1", 500, "d", "descript")
		if err != nil {
			t.Fatalf("Failed to update balance: %v", err)
		}

		// assert updated balance
		assertBalance(t, 1000.0, 0)
		// assert that the transaction was created
		assertTransactionCreated(t, "1", 500, "d", "descript")
	})
}

func prepareDb() {
	helpers.InitTestDB()
	createAccount()
}

func createAccount() {
	account := models.Account{Name: "test account", LimitAmount: 1000}
	db.DB.Create(&account)
	balance := models.Balance{AccountID: account.ID, Amount: 500}
	db.DB.Create(&balance)
}

func assertBalance(
	t *testing.T,
	expectedLimitAmount int,
	expectedAmount int,
) {
	balance, err := queries.GetBalance("1")
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}

	if balance.LimitAmount != expectedLimitAmount {
		t.Errorf("Expected LimitAmount to be %d, got %d", expectedLimitAmount, balance.LimitAmount)
	}
	if balance.Amount != expectedAmount {
		t.Errorf("Expected Amount to be %d, got %d", expectedAmount, balance.Amount)
	}
}

func assertTransactionCreated(t *testing.T, accountId string, amount int, transactionType string, description string) {
	var transaction models.Transaction
	if err := db.DB.Where("account_id = ? AND amount = ? AND transaction_type = ?", accountId, amount, transactionType).First(&transaction).Error; err != nil {
		t.Fatalf("Failed to find transaction: %v", err)
	}

	if transaction.AccountID != uint(parseUint(accountId)) {
		t.Errorf("Expected AccountID to be %v, got %v", accountId, transaction.AccountID)
	}
	if transaction.Amount != amount {
		t.Errorf("Expected Amount to be %v, got %v", amount, transaction.Amount)
	}
	if transaction.TransactionType != transactionType {
		t.Errorf("Expected TransactionType to be %v, got %v", transactionType, transaction.TransactionType)
	}
	if transaction.Description != description {
		t.Errorf("Expected Description to be %v, got %v", description, transaction.Description)
	}
	if time.Since(transaction.Date) > time.Minute {
		t.Errorf("Transaction Date is not recent")
	}
}
