package internal

import (
	"github.com/dmedinao1/go-web-practica/pkg/store"
	"time"
)

type Transaction struct {
	Id              int       `json:"id"`
	TransactionCode string    `json:"transactionCode"`
	Currency        string    `json:"currency"`
	Quantity        float64   `json:"quantity"`
	Transmitter     string    `json:"transmitter"`
	TransactionDate time.Time `json:"transactionDate"`
}

type TransactionRepository interface {
	FindAll() ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	GetLastId() (int, error)
	GetById(id int) (*Transaction, error)
	Replace(id int, transaction Transaction) (Transaction, error)
	DeleteById(id int) error
}

func GetTransactionRepository(store store.Store) TransactionRepository {
	return &transactionRepository{store}
}

type transactionRepository struct {
	store store.Store
}

func (t *transactionRepository) FindAll() ([]Transaction, error) {
	var transactions []Transaction

	err := t.store.Read(&transactions)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionRepository) Save(transaction Transaction) (Transaction, error) {
	transactions, err := t.FindAll()

	if err != nil {
		return Transaction{}, err
	}

	transactions = append(transactions, transaction)

	err = t.store.Write(&transactions)

	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionRepository) GetLastId() (int, error) {
	transactions, err := t.FindAll()

	if err != nil {
		return 0, err
	}

	if len(transactions) == 0 {
		return 0, nil
	}

	var lastId int

	for i, transaction := range transactions {
		if i == 0 || lastId < transaction.Id {
			lastId = transaction.Id
		}
	}

	return lastId, nil
}

func (t *transactionRepository) GetById(id int) (*Transaction, error) {
	transactions, err := t.FindAll()

	if err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		if transaction.Id == id {
			return &transaction, nil
		}
	}

	return nil, nil
}

func (t *transactionRepository) Replace(id int, transaction Transaction) (Transaction, error) {
	transactions, err := t.FindAll()

	if err != nil {
		return Transaction{}, err
	}

	var transactionIndex int
	for i, currentTransaction := range transactions {
		if currentTransaction.Id == id {
			transactionIndex = i
			break
		}
	}

	transactions[transactionIndex] = transaction

	return transaction, nil
}

func (t *transactionRepository) DeleteById(id int) error {
	transactions, err := t.FindAll()

	if err != nil {
		return err
	}

	var transactionIndex int
	for i, transaction := range transactions {
		if transaction.Id == id {
			transactionIndex = i
			break
		}
	}

	transactions = append(transactions[:transactionIndex], transactions[transactionIndex+1:]...)

	err = t.store.Write(&transactions)

	if err != nil {
		return err
	}

	return nil
}
