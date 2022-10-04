package models

import (
	"encoding/json"
	"os"
	"time"
)

type Transaction struct {
	Id              int       `json:"id"`
	TransactionCode string    `json:"transactionCode" binding:"required"`
	Currency        string    `json:"currency" binding:"required"`
	Quantity        float64   `json:"quantity" binding:"required"`
	Transmitter     string    `json:"transmitter" binding:"required"`
	TransactionDate time.Time `json:"transactionDate" binding:"required"`
}

func (t Transaction) GetFromFile(filePath string) (*[]Transaction, error) {
	var transactions []Transaction

	rawData, err := os.ReadFile(filePath)

	if err != nil {
		return &transactions, err
	}

	err = json.Unmarshal(rawData, &transactions)

	if err != nil {
		return &transactions, err
	}

	return &transactions, nil
}

func (t Transaction) CheckFilters(params *Transaction) bool {
	if params.Id != 0 && t.Id != params.Id {
		return false
	}

	if params.TransactionCode != "" && t.TransactionCode != params.TransactionCode {
		return false
	}

	if params.Currency != "" && t.Currency != params.Currency {
		return false
	}

	if params.Quantity != 0 && t.Quantity != params.Quantity {
		return false
	}

	if params.Transmitter != "" && t.Transmitter != params.Transmitter {
		return false
	}

	if !params.TransactionDate.IsZero() && t.TransactionDate.Equal(params.TransactionDate) {
		return false
	}

	return true
}
