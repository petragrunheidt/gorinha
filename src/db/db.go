package db

import (
	"context"
	"fmt"
	"log"
	
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
