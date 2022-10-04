package services

import (
	"github.com/dmedinao1/ejercicio-TM-04-10/models"
	"github.com/dmedinao1/ejercicio-TM-04-10/repositories"
)

type TransactionService struct {
	transactionRepository *repositories.TransactionRepository
}

var transactionService = TransactionService{transactionRepository: repositories.GetTransactionRepository()}

func GetTransactionService() *TransactionService {
	return &transactionService
}

func (t TransactionService) SaveTransaction(transaction *models.Transaction) *models.Transaction {
	return t.transactionRepository.SaveTransaction(transaction)
}

func (t TransactionService) GetAll() []models.Transaction {
	return t.transactionRepository.FindAll()
}
