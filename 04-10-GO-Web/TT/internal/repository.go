package internal

import "time"

type Transaction struct {
	Id              int       `json:"id"`
	TransactionCode string    `json:"transactionCode"`
	Currency        string    `json:"currency"`
	Quantity        float64   `json:"quantity"`
	Transmitter     string    `json:"transmitter"`
	TransactionDate time.Time `json:"transactionDate"`
}

type TransactionRepository interface {
	FindAll() []Transaction
	Save(transaction Transaction) Transaction
	GetLastId() int
}

func GetTransactionRepository() TransactionRepository {
	return transactionRepository{transactions: &[]Transaction{}}
}

type transactionRepository struct {
	transactions *[]Transaction
}

func (t transactionRepository) FindAll() []Transaction {
	return *t.transactions
}

func (t transactionRepository) Save(transaction Transaction) Transaction {
	*(t.transactions) = append(*(t.transactions), transaction)
	return transaction
}

func (t transactionRepository) GetLastId() int {
	transactions := *t.transactions
	if len(transactions) == 0 {
		return 0
	}

	return transactions[len(transactions)-1].Id
}
