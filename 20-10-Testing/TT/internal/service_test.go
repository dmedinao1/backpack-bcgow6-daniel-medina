package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type storeServiceStub struct {
	WasReadCalled    bool
	WasWriteCalled   bool
	FakeTransactions []Transaction
}

func (s *storeServiceStub) Read(data interface{}) error {
	s.WasReadCalled = true

	toReadPointer, ok := data.(*[]Transaction)

	if !ok {
		return errors.New("type conversion fails :(")
	}

	*toReadPointer = s.FakeTransactions

	return nil
}

func (s *storeServiceStub) Write(data interface{}) error {
	s.WasWriteCalled = true

	toWritePointer, ok := data.(*[]Transaction)

	if !ok {
		return errors.New("type conversion fails :(")
	}

	s.FakeTransactions = *toWritePointer

	return nil
}

func Test_transactionService_UpdateCodeAndQuantityById(t1 *testing.T) {
	transactions := []Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
	}

	const ID = 1
	const QUANTITY float64 = 500
	const CODE = "TC2"

	storeStub := storeServiceStub{FakeTransactions: transactions}

	repository := GetTransactionRepository(&storeStub)
	service := GetTransactionService(repository)

	err := service.UpdateCodeAndQuantityById(ID, CODE, QUANTITY)

	assert.NoError(t1, err, "Got an error when calling UpdateCodeAndQuantityById | %s", err)

	updated, _ := repository.GetById(ID)

	assert.True(t1, storeStub.WasReadCalled)
	assert.True(t1, storeStub.WasWriteCalled)
	assert.Equal(t1, QUANTITY, updated.Quantity)
	assert.Equal(t1, CODE, updated.TransactionCode)
}

func Test_transactionService_DeleteById(t1 *testing.T) {
	transactions := []Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
	}

	const ID = 1

	storeStub := storeServiceStub{FakeTransactions: transactions}

	repository := GetTransactionRepository(&storeStub)
	service := GetTransactionService(repository)

	err := service.DeleteById(ID)

	assert.NoError(t1, err, "Got an error when calling UpdateCodeAndQuantityById | %s", err)

	_, err = repository.GetById(ID)

	assert.Equal(t1, len(storeStub.FakeTransactions), len(transactions)-1)
	assert.Error(t1, err, "Didn't get an error when calling get by id of deleted element")

	const IdNotExists = 2

	err = service.DeleteById(IdNotExists)

	assert.Error(t1, err, "Didn't get an error when calling delete by id when id does not exists")
}
