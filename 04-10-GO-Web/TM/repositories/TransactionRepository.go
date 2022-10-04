package repositories

import "github.com/dmedinao1/ejercicio-TM-04-10/models"

type TransactionRepository struct {
	transactions []models.Transaction
}

var transactionRepository = TransactionRepository{transactions: []models.Transaction{}}

func GetTransactionRepository() *TransactionRepository {
	return &transactionRepository
}

func (t *TransactionRepository) SaveTransaction(transaction *models.Transaction) *models.Transaction {
	id := t.getLastId() + 1

	transaction.Id = id

	t.transactions = append(t.transactions, *transaction)

	return transaction
}

func (t *TransactionRepository) getLastId() int {
	if len(t.transactions) == 0 {
		return 0
	}

	return t.transactions[len(t.transactions)-1].Id
}

func (t *TransactionRepository) FindAll() []models.Transaction {
	return t.transactions
}
