package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type storeStub struct {
	WasReadCalled    bool
	WasWriteCalled   bool
	FakeTransactions []Transaction
}

func (s *storeStub) Read(data interface{}) error {
	s.WasReadCalled = true

	toReadPointer, ok := data.(*[]Transaction)

	if !ok {
		return errors.New("type conversion fails :(")
	}

	*toReadPointer = s.FakeTransactions

	return nil
}

func (s *storeStub) Write(data interface{}) error {
	s.WasWriteCalled = true

	toWritePointer, ok := data.(*[]Transaction)

	if !ok {
		return errors.New("type conversion fails :(")
	}

	s.FakeTransactions = *toWritePointer

	return nil
}

func Test_transactionRepository_FindAll(t1 *testing.T) {
	transactions := []Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
		{2, "TC2", "COP", 134.1, "Central Bank", time.Now()},
	}

	storeStub := storeStub{FakeTransactions: transactions}

	repository := GetTransactionRepository(&storeStub)

	transactions, err := repository.FindAll()

	assert.NoError(t1, err, "Expect not get an error but fails | got error: '%s'", err)
	assert.True(t1, storeStub.WasReadCalled)
	assert.Equal(t1, transactions, transactions)
}

func Test_transactionRepository_Replace(t1 *testing.T) {
	transactions := []Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
		{2, "TC2", "COP", 134.1, "Central Bank", time.Now()},
	}

	const ID = 1
	const QUANTITY float64 = 500
	const TRANSMITTER = "Another bank"

	storeStub := storeStub{FakeTransactions: transactions}

	repository := GetTransactionRepository(&storeStub)

	toReplace, _ := repository.GetById(ID)

	toReplace.Quantity = QUANTITY
	toReplace.Transmitter = TRANSMITTER

	_, err := repository.Replace(ID, *toReplace)

	assert.NoError(t1, err, "Got an error when calling replace | %s", err)

	replaced, _ := repository.GetById(ID)

	assert.Equal(t1, QUANTITY, replaced.Quantity)
	assert.Equal(t1, TRANSMITTER, replaced.Transmitter)
}
