package internal

import (
	"fmt"
	"net/http"
	"time"
)

type ApiError struct {
	Code         int
	ErrorMessage string
}

func (a ApiError) Error() string {
	return a.ErrorMessage
}

type TransactionService interface {
	FindAll() []Transaction

	SaveTransaction(
		transactionCode string,
		currency string,
		quantity float64,
		transmitter string,
		transactionDate time.Time,
	) Transaction

	ReplaceTransaction(
		id int,
		code string,
		currency string,
		quantity float64,
		transmitter string,
		date time.Time,
	) (Transaction, error)
	UpdateCodeAndQuantityById(id int, transactionCode string, quantity float64) error
	DeleteById(id int) error
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

func (t transactionService) ReplaceTransaction(id int, code string, currency string, quantity float64, transmitter string, date time.Time) (Transaction, error) {

	foundedTransaction := t.transactionRepository.GetById(id)

	if foundedTransaction == nil {
		return Transaction{}, ApiError{Code: http.StatusNotFound, ErrorMessage: fmt.Sprintf("Id %v no encontrado", id)}
	}

	toReplace := Transaction{
		Id:              id,
		TransactionCode: code,
		Currency:        currency,
		Quantity:        quantity,
		Transmitter:     transmitter,
		TransactionDate: date,
	}

	return t.transactionRepository.Replace(id, toReplace), nil
}

func (t transactionService) UpdateCodeAndQuantityById(id int, transactionCode string, quantity float64) error {
	foundedTransaction := t.transactionRepository.GetById(id)

	if foundedTransaction == nil {
		return ApiError{Code: http.StatusNotFound, ErrorMessage: fmt.Sprintf("Id %v no encontrado", id)}
	}

	(*foundedTransaction).TransactionCode = transactionCode
	(*foundedTransaction).Quantity = quantity

	t.transactionRepository.Replace(id, *foundedTransaction)

	return nil
}

func (t transactionService) DeleteById(id int) error {
	foundedTransaction := t.transactionRepository.GetById(id)

	if foundedTransaction == nil {
		return ApiError{Code: http.StatusNotFound, ErrorMessage: fmt.Sprintf("Id %v no encontrado", id)}
	}

	t.transactionRepository.DeleteById(id)

	return nil
}
