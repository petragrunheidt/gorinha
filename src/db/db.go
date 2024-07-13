package db

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"gorinha/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
)

var Gorm *gorm.DB

func Init() {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
			log.Fatalf("error determining current file path")
	}

	configPath := filepath.Join(filepath.Dir(currentFile), "./config.yml")

	config, err := LoadConfig(configPath)
	if err != nil {
			log.Fatalf("Error loading config file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.Host, config.User, config.Password, config.Name, config.Port)

	Gorm, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Database connection initialized")
}

func Close() {
	sqlDB, err := Gorm.DB()
	if err != nil {
			log.Fatalf("Error getting SQL DB: %v", err)
	}
	sqlDB.Close()
	fmt.Println("Database connection closed")
}

func Migrate() {
	err := Gorm.AutoMigrate(
		&models.Account{},
		&models.Balance{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("Database migration completed")
}

func Drop() {
	// Get all table names
	var tables []string
	Gorm.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)

	// Drop each table
	for _, table := range tables {
		Gorm.Migrator().DropTable(table)
	}
}

