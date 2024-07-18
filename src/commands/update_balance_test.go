package commands

import (
	"testing"

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

		err := UpdateBalance("1", 500, "c")
		if err != nil {
			t.Fatalf("Failed to update balance: %v", err)
		}

		// assert updated balance
		assertBalance(t, 500.0, 500.0)
	})
	t.Run("Debit", func(t *testing.T) {
		prepareDb()

		// assert initial balance
		assertBalance(t, 1000.0, 500.0)

		err := UpdateBalance("1", 500, "d")
		if err != nil {
			t.Fatalf("Failed to update balance: %v", err)
		}

		// assert updated balance
		assertBalance(t, 1000.0, 0)
	})
}

func prepareDb() {
	helpers.InitTestDB()
	createAccount()
}

func createAccount() {
	account := models.Account{Name: "test account", LimitAmount: 1000}
	db.Gorm.Create(&account)
	balance := models.Balance{AccountID: account.ID, Amount: 500}
	db.Gorm.Create(&balance)
}

func assertBalance(
	t *testing.T,
	expectedLimitAmount float64,
	expectedAmount float64,
) {
	balance, err := queries.GetBalance("1")
	if err != nil {
		t.Fatalf("Failed to get balance: %v", err)
	}

	if balance.LimitAmount != expectedLimitAmount {
		t.Errorf("Expected LimitAmount to be %.2f, got %.2f", expectedLimitAmount, balance.LimitAmount)
	}
	if balance.Amount != expectedAmount {
		t.Errorf("Expected Amount to be %.2f, got %.2f", expectedAmount, balance.Amount)
	}
}
