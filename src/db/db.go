package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool

func Init() {
	config, err := LoadConfig("src/db/config.yml")

	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.DB.User, config.DB.Password,
		config.DB.Host, config.DB.Port,
		config.DB.Name)

	DBPool, err = pgxpool.New(context.Background(), dbURL)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Database connection pool initialized")
}

func Close() {
	DBPool.Close()
	fmt.Println("Database connection pool closed")
}

func PingDB() error {
	log.Println("doing ping...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DBPool.Ping(ctx); err != nil {
		log.Printf("Ping bad :(, %s", err)
		return fmt.Errorf("unable to ping database: %w", err)
	}
	log.Println("good ping it work")
	return nil
}

func QueryExample() ([]byte, error) {
	sqlStatement := "SELECT a.limit_amount, b.amount FROM accounts AS a JOIN balances AS b ON a.id = b.account_id WHERE a.id = 1"

	rows, err := DBPool.Query(context.Background(), sqlStatement)
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	// Define a struct to hold the query results
	type Result struct {
			LimitAmount float64 `json:"limit_amount"`
			Amount      float64 `json:"amount"`
	}

	var results []Result

	// Iterate through the result set
	for rows.Next() {
			var result Result
			if err := rows.Scan(&result.LimitAmount, &result.Amount); err != nil {
					return nil, err
			}
			results = append(results, result)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
			return nil, err
	}

	// Marshal results to JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
			return nil, err
	}

	return jsonData, nil
}
