package internal

import (
	"errors"
	"github.com/dmedinao1/go-web-practica/internal/domain"
	"github.com/dmedinao1/go-web-practica/pkg/store"
)

type TransactionRepository interface {
	FindAll() ([]domain.Transaction, error)
	Save(transaction domain.Transaction) (domain.Transaction, error)
	GetLastId() (int, error)
	GetById(id int) (*domain.Transaction, error)
	Replace(id int, transaction domain.Transaction) (domain.Transaction, error)
	DeleteById(id int) error
}

func GetTransactionRepository(store store.Store) TransactionRepository {
	return &transactionRepository{store}
}

type transactionRepository struct {
	store store.Store
}

func (t *transactionRepository) FindAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	err := t.store.Read(&transactions)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionRepository) Save(transaction domain.Transaction) (domain.Transaction, error) {
	transactions, err := t.FindAll()

	if err != nil {
		return domain.Transaction{}, err
	}

	transactions = append(transactions, transaction)

	err = t.store.Write(&transactions)

	if err != nil {
		return domain.Transaction{}, err
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

func (t *transactionRepository) GetById(id int) (*domain.Transaction, error) {
	transactions, err := t.FindAll()

	if err != nil {
		return nil, err
	}

	for _, transaction := range transactions {
		if transaction.Id == id {
			return &transaction, nil
		}
	}

	return nil, errors.New("not found")
}

func (t *transactionRepository) Replace(id int, transaction domain.Transaction) (domain.Transaction, error) {
	transactions, err := t.FindAll()

	if err != nil {
		return domain.Transaction{}, err
	}

	var transactionIndex int
	for i, currentTransaction := range transactions {
		if currentTransaction.Id == id {
			transactionIndex = i
			break
		}
	}

	transactions[transactionIndex] = transaction

	err = t.store.Write(&transactions)

	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionRepository) DeleteById(id int) error {
	transactions, err := t.FindAll()

	if err != nil {
		return err
	}

	var transactionIndex = -1
	for i, transaction := range transactions {
		if transaction.Id == id {
			transactionIndex = i
			break
		}
	}

	if transactionIndex == -1 {
		return errors.New("not found")
	}

	transactions = append(transactions[:transactionIndex], transactions[transactionIndex+1:]...)

	err = t.store.Write(&transactions)

	if err != nil {
		return err
	}

	return nil
}
