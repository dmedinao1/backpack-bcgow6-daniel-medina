package mocks

import (
	"errors"
	"github.com/dmedinao1/go-web-practica/internal/domain"
	"github.com/dmedinao1/go-web-practica/pkg/store"
)

func GetFunStoreStub(transactions []domain.Transaction) store.Store {
	return &storeServiceStub{FakeTransactions: transactions}
}

func GetStoreStubWithErrors() store.Store {
	return &storeErrorMock{}
}

type storeServiceStub struct {
	WasReadCalled    bool
	WasWriteCalled   bool
	FakeTransactions []domain.Transaction
}

func (s *storeServiceStub) Read(data interface{}) error {
	s.WasReadCalled = true

	toReadPointer, ok := data.(*[]domain.Transaction)

	if !ok {
		return errors.New("type conversion fails :(")
	}

	*toReadPointer = s.FakeTransactions

	return nil
}

func (s *storeServiceStub) Write(data interface{}) error {
	s.WasWriteCalled = true

	toWritePointer, ok := data.(*[]domain.Transaction)

	if !ok {
		return errors.New("type conversion fails :(")
	}

	s.FakeTransactions = *toWritePointer

	return nil
}

type storeErrorMock struct {
}

func (s *storeErrorMock) Read(data interface{}) error {
	return errors.New("error in read method")
}

func (s *storeErrorMock) Write(data interface{}) error {
	return errors.New("error in write method")
}
