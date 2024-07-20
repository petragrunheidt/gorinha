package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gorinha/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	
	runSeeds()
}

func runSeeds() {
	isLeader := os.Getenv("LEADER") == "true"
	isDebug := gin.Mode() == "debug"
	
	if isLeader && isDebug {
		Drop()
		Migrate()
		if err := seedData(Gorm); err != nil {
			fmt.Printf("Error seeding data: %v\n", err)
		} else {
			fmt.Println("Development database populated")
		}
	}
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

func seedData(db *gorm.DB) error {
	// Seed accounts
	accounts := []models.Account{
		{Name: "o barato sai caro", LimitAmount: 1000 * 100},
		{Name: "zan corp ltda", LimitAmount: 800 * 100},
		{Name: "les cruders", LimitAmount: 10000 * 100},
		{Name: "padaria joia de cocaia", LimitAmount: 100000 * 100},
		{Name: "kid mais", LimitAmount: 5000 * 100},
	}

	for _, account := range accounts {
		if err := db.Create(&account).Error; err != nil {
			return err
		}
	}

	// Seed balances
	var accountIDs []uint
	db.Model(&models.Account{}).Pluck("id", &accountIDs)

	balances := make([]models.Balance, len(accountIDs))
	for i, id := range accountIDs {
		balances[i] = models.Balance{AccountID: id, Amount: 0}
	}

	if err := db.Create(&balances).Error; err != nil {
		return err
	}

	return nil
}
