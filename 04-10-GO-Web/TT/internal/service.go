package internal

import "time"

type TransactionService interface {
	FindAll() []Transaction
	SaveTransaction(
		transactionCode string,
		currency string,
		quantity float64,
		transmitter string,
		transactionDate time.Time,
	) Transaction
}

type transactionService struct {
	transactionRepository TransactionRepository
}

func GetTransactionService(repository TransactionRepository) TransactionService {
	return transactionService{transactionRepository: repository}
}

func (t transactionService) FindAll() []Transaction {
	return t.transactionRepository.FindAll()
}

func (t transactionService) SaveTransaction(transactionCode string, currency string, quantity float64, transmitter string, transactionDate time.Time) Transaction {
	id := t.transactionRepository.GetLastId() + 1
	toSave := Transaction{
		Id:              id,
		TransactionCode: transactionCode,
		Currency:        currency,
		Quantity:        quantity,
		Transmitter:     transmitter,
		TransactionDate: transactionDate,
	}

	return t.transactionRepository.Save(toSave)
}