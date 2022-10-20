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
	FindAll() ([]Transaction, error)
	SaveTransaction(transactionCode string, currency string, quantity float64, transmitter string, transactionDate time.Time) (Transaction, error)
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

func (t transactionService) FindAll() ([]Transaction, error) {
	transactions, err := t.transactionRepository.FindAll()

	if err != nil {
		return nil, ApiError{Code: http.StatusInternalServerError, ErrorMessage: "Error en la base de datos | " + err.Error()}
	}
	return transactions, nil
}

func (t transactionService) SaveTransaction(transactionCode string, currency string, quantity float64, transmitter string, transactionDate time.Time) (Transaction, error) {
	lastId, err := t.transactionRepository.GetLastId()

	if err != nil {
		return Transaction{}, err
	}

	id := lastId + 1

	toSave := Transaction{
		Id:              id,
		TransactionCode: transactionCode,
		Currency:        currency,
		Quantity:        quantity,
		Transmitter:     transmitter,
		TransactionDate: transactionDate,
	}

	saved, err := t.transactionRepository.Save(toSave)

	if err != nil {
		return Transaction{}, ApiError{Code: http.StatusInternalServerError, ErrorMessage: "Error en la base de datos | " + err.Error()}
	}

	return saved, nil
}

func (t transactionService) ReplaceTransaction(id int, code string, currency string, quantity float64, transmitter string, date time.Time) (Transaction, error) {

	foundedTransaction, _ := t.transactionRepository.GetById(id)

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

	saved, err := t.transactionRepository.Replace(id, toReplace)

	if err != nil {
		return Transaction{}, ApiError{Code: http.StatusInternalServerError, ErrorMessage: "Error en la base de datos | " + err.Error()}
	}

	return saved, nil
}

func (t transactionService) UpdateCodeAndQuantityById(id int, transactionCode string, quantity float64) error {
	foundedTransaction, err := t.transactionRepository.GetById(id)

	if err != nil {
		return ApiError{Code: http.StatusInternalServerError, ErrorMessage: "Error en la base de datos | " + err.Error()}
	}

	if foundedTransaction == nil {
		return ApiError{Code: http.StatusNotFound, ErrorMessage: fmt.Sprintf("Id %v no encontrado", id)}
	}

	foundedTransaction.TransactionCode = transactionCode
	foundedTransaction.Quantity = quantity

	_, err = t.transactionRepository.Replace(id, *foundedTransaction)
	if err != nil {
		return ApiError{Code: http.StatusInternalServerError, ErrorMessage: "Error en la base de datos | " + err.Error()}
	}

	return nil
}

func (t transactionService) DeleteById(id int) error {
	foundedTransaction, _ := t.transactionRepository.GetById(id)

	if foundedTransaction == nil {
		return ApiError{Code: http.StatusNotFound, ErrorMessage: fmt.Sprintf("Id %v no encontrado", id)}
	}

	err := t.transactionRepository.DeleteById(id)

	if err != nil {
		return ApiError{Code: http.StatusInternalServerError, ErrorMessage: "Error en la base de datos | " + err.Error()}
	}

	return nil
}
