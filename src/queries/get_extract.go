package queries

import (
	"fmt"
	"time"

	"gorinha/src/db"
	"gorinha/src/models"
)

type Extract struct {
	ExtractBalance     ExtractBalance      `json:"saldo"`
	TransactionRecords []TransactionRecord `json:"ultimas_transacoes"`
}

type ExtractBalance struct {
	Amount      float64   `json:"total"`
	Date        time.Time `json:"data_extrato"`
	LimitAmount float64   `json:"limite"`
}

type TransactionRecord struct {
	Amount      float64   `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	Timestamp   time.Time `json:"realizada_em"`
}

func GetExtract(id string) (Extract, error) {
	var extract Extract

	balance, err := GetBalance(id)
	if err != nil {
		return extract, err
	}

	extract.ExtractBalance = ExtractBalance{
		Amount:      balance.Amount,
		Date:        time.Now(),
		LimitAmount: balance.LimitAmount,
	}

	extract.TransactionRecords, err = getLast10Extracts(id)

	fmt.Printf("Extract: %+v\n", extract)
	return extract, err
}

func getLast10Extracts(id string) ([]TransactionRecord, error) {
	var transactions []models.Transaction
	var transactionRecords []TransactionRecord

	err := db.Gorm.Where("account_id = ?", id).
		Order("id DESC").
		Limit(10).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	for _, t := range transactions {
		record := TransactionRecord{
			Amount:      t.Amount,
			Type:        t.TransactionType,
			Description: t.Description,
			Timestamp:   t.Date,
		}
		transactionRecords = append(transactionRecords, record)
	}

	return transactionRecords, nil
}
