package queries

import (
	"testing"

	"gorinha/src/db"
	"gorinha/src/helpers"
	"gorinha/src/models"
)

func TestGetBalance(t *testing.T) {
	helpers.InitTestDB()
	createAccount()

	balance, err := GetBalance("1")
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}

	if balance.LimitAmount != 1000.0 {
		t.Errorf("Expected LimitAmount to be 1000.0, got %v", balance.LimitAmount)
	}
	if balance.Amount != 500.0 {
		t.Errorf("Expected Amount to be 500.0, got %v", balance.Amount)
	}

	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}

func createAccount() {
	account := models.Account{Name: "test account", LimitAmount: 1000}
	db.DB.Create(&account)
	balance := models.Balance{AccountID: account.ID, Amount: 500}
	db.DB.Create(&balance)
}
