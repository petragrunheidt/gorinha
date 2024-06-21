package queries

import (
	"context"
	"fmt"
	"gorinha/src/db"
	"log"
)

type Balance struct {
	LimitAmount float64 `json:"limit_amount"`
	Amount      float64 `json:"amount"`
}

func GetBalance(id string) ([]Balance, error) {
	sqlStatement := `
	SELECT a.limit_amount, b.amount 
	FROM accounts AS a 
	JOIN balances AS b ON a.id = b.account_id 
	WHERE a.id = $1
	`
	log.Printf("\n\n\n\n\n\nstarting query with id: %s", id)
	rows, err := db.DBPool.Query(context.Background(), sqlStatement, id)
	if err != nil {
		log.Fatalf("Error querying database: %v\n", err)
	}
	defer rows.Close()

	var balances []Balance

	for rows.Next() {
		var balance Balance

		err := rows.Scan(&balance.LimitAmount, &balance.Amount)
		log.Printf("added balance: %v, %v", balance.LimitAmount, balance.Amount)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		balances = append(balances, balance)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating rows: %v\n", err)
	}

	defer db.DBPool.Close()
	fmt.Println("Database connection pool closed")
	return balances, err
}
