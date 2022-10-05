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
	GetById(id int) *Transaction
	Replace(id int, transaction Transaction) Transaction
	DeleteById(id int)
}

func GetTransactionRepository() TransactionRepository {
	return &transactionRepository{transactions: &[]Transaction{}}
}

type transactionRepository struct {
	transactions *[]Transaction
}

func (t *transactionRepository) FindAll() []Transaction {
	return *t.transactions
}

func (t *transactionRepository) Save(transaction Transaction) Transaction {
	*(t.transactions) = append(*(t.transactions), transaction)
	return transaction
}

func (t *transactionRepository) GetLastId() int {
	transactions := *t.transactions
	if len(transactions) == 0 {
		return 0
	}

	return transactions[len(transactions)-1].Id
}

func (t *transactionRepository) GetById(id int) *Transaction {
	for _, transaction := range *t.transactions {
		if transaction.Id == id {
			return &transaction
		}
	}

	return nil
}

func (t *transactionRepository) Replace(id int, toReplace Transaction) Transaction {
	var transactionIndex int
	for i, transaction := range *t.transactions {
		if transaction.Id == id {
			transactionIndex = i
			break
		}
	}
	transactions := *(t.transactions)
	transactions[transactionIndex] = toReplace

	return toReplace
}

func (t *transactionRepository) DeleteById(id int) {
	var transactionIndex int
	for i, transaction := range *t.transactions {
		if transaction.Id == id {
			transactionIndex = i
			break
		}
	}

	transactions := *t.transactions

	transactions = append(transactions[:transactionIndex], transactions[transactionIndex+1:]...)

	t.transactions = &transactions
}
