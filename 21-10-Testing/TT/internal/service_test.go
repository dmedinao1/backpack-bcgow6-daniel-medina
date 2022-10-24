package internal

import (
	"errors"
	"fmt"
	"github.com/dmedinao1/go-web-practica/internal/domain"
	"github.com/dmedinao1/go-web-practica/internal/mocks"
	"github.com/dmedinao1/go-web-practica/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func GetRepositoryMock(store store.Store, args RepositoryMockArgs) TransactionRepository {
	repository := GetTransactionRepository(store)

	return &repositoryMock{
		args: args,
		repo: repository,
	}
}

type RepositoryMockArgs struct {
	HaveErrorInFindAll    bool
	HaveErrorInSave       bool
	HaveErrorInGetLastId  bool
	HaveErrorInGetById    bool
	HaveErrorInReplace    bool
	HaveErrorInDeleteById bool
}

type repositoryMock struct {
	args RepositoryMockArgs
	repo TransactionRepository
}

func (r *repositoryMock) FindAll() ([]domain.Transaction, error) {
	if r.args.HaveErrorInFindAll {
		return nil, errors.New("error in FindAll")
	}
	return r.repo.FindAll()
}

func (r *repositoryMock) Save(transaction domain.Transaction) (domain.Transaction, error) {
	if r.args.HaveErrorInSave {
		return domain.Transaction{}, errors.New("error in Save")
	}
	return r.repo.Save(transaction)
}

func (r *repositoryMock) GetLastId() (int, error) {
	if r.args.HaveErrorInGetLastId {
		return 0, errors.New("error in GetLastId")
	}
	return r.repo.GetLastId()
}

func (r *repositoryMock) GetById(id int) (*domain.Transaction, error) {
	if r.args.HaveErrorInGetById {
		return nil, errors.New("get by id error")
	}

	return r.repo.GetById(id)
}

func (r *repositoryMock) Replace(id int, transaction domain.Transaction) (domain.Transaction, error) {
	if r.args.HaveErrorInReplace {
		return domain.Transaction{}, errors.New("get by id error")
	}

	return r.repo.Replace(id, transaction)
}

func (r *repositoryMock) DeleteById(id int) error {
	if r.args.HaveErrorInDeleteById {
		return errors.New("error in DeleteById")
	}
	return r.repo.DeleteById(id)
}

func Test_transactionService_UpdateCodeAndQuantityById(t1 *testing.T) {
	transactions := []domain.Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
	}

	const ID = 1
	const QUANTITY float64 = 500
	const CODE = "TC2"

	storeStub := mocks.GetFunStoreStub(transactions)

	repository := GetTransactionRepository(storeStub)
	service := GetTransactionService(repository)

	err := service.UpdateCodeAndQuantityById(ID, CODE, QUANTITY)

	assert.NoError(t1, err, "Got an error when calling UpdateCodeAndQuantityById | %s", err)

	updated, _ := repository.GetById(ID)

	assert.Equal(t1, QUANTITY, updated.Quantity)
	assert.Equal(t1, CODE, updated.TransactionCode)
}

func Test_transactionService_DeleteById(t1 *testing.T) {
	transactions := []domain.Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
	}

	const ID = 1

	storeStub := mocks.GetFunStoreStub(transactions)

	repository := GetTransactionRepository(storeStub)
	service := GetTransactionService(repository)

	err := service.DeleteById(ID)

	assert.NoError(t1, err, "Got an error when calling UpdateCodeAndQuantityById | %s", err)

	_, err = repository.GetById(ID)

	assert.Error(t1, err, "Didn't get an error when calling get by id of deleted element")

	const IdNotExists = 2

	err = service.DeleteById(IdNotExists)

	assert.Error(t1, err, "Didn't get an error when calling delete by id when id does not exists")
}

func Test_transactionService_FindAll(t1 *testing.T) {
	transactions := []domain.Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
	}

	funStore := mocks.GetFunStoreStub(transactions)
	errorStore := mocks.GetStoreStubWithErrors()

	funRepository := GetTransactionRepository(funStore)
	errorRepository := GetTransactionRepository(errorStore)

	type fields struct {
		transactionRepository TransactionRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Transaction
		wantErr assert.ErrorAssertionFunc
	}{
		{"Fun case for get all", fields{funRepository}, transactions, assert.NoError},
		{"Error case for get all", fields{errorRepository}, nil, assert.Error},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := GetTransactionService(tt.fields.transactionRepository)
			got, err := t.FindAll()
			if !tt.wantErr(t1, err, fmt.Sprintf("FindAll()")) {
				return
			}
			assert.Equalf(t1, tt.want, got, "FindAll()")
		})
	}
}

func Test_transactionService_SaveTransaction(t1 *testing.T) {
	var transactions []domain.Transaction

	fakeTransaction := domain.Transaction{
		TransactionCode: "TC",
		Currency:        "COP",
		Quantity:        500,
		Transmitter:     "Central bank",
		TransactionDate: time.Time{},
	}

	funStore := mocks.GetFunStoreStub(transactions)
	errorStore := mocks.GetStoreStubWithErrors()

	funRepository := GetTransactionRepository(funStore)
	errorRepository := GetTransactionRepository(errorStore)

	type fields struct {
		transactionRepository TransactionRepository
	}
	type args struct {
		transactionCode string
		currency        string
		quantity        float64
		transmitter     string
		transactionDate time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Transaction
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Fun case 1 for save transaction",
			fields{funRepository},
			args{fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{Id: 1, TransactionCode: fakeTransaction.TransactionCode, Currency: fakeTransaction.Currency, Quantity: fakeTransaction.Quantity, Transmitter: fakeTransaction.Transmitter, TransactionDate: fakeTransaction.TransactionDate},
			assert.NoError,
		},
		{
			"Error case for save transaction",
			fields{errorRepository},
			args{fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{},
			assert.Error,
		},
		{
			"Fun case 2 for save transaction (check id sequence)",
			fields{funRepository},
			args{fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{Id: 2, TransactionCode: fakeTransaction.TransactionCode, Currency: fakeTransaction.Currency, Quantity: fakeTransaction.Quantity, Transmitter: fakeTransaction.Transmitter, TransactionDate: fakeTransaction.TransactionDate},
			assert.NoError,
		},
		{
			"Error case for save transaction (error in repository)",
			fields{GetRepositoryMock(funStore, RepositoryMockArgs{HaveErrorInSave: true})},
			args{fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := transactionService{
				transactionRepository: tt.fields.transactionRepository,
			}
			got, err := t.SaveTransaction(tt.args.transactionCode, tt.args.currency, tt.args.quantity, tt.args.transmitter, tt.args.transactionDate)
			if !tt.wantErr(t1, err, fmt.Sprintf("SaveTransaction(%v, %v, %v, %v, %v)", tt.args.transactionCode, tt.args.currency, tt.args.quantity, tt.args.transmitter, tt.args.transactionDate)) {
				return
			}
			assert.Equalf(t1, tt.want, got, "SaveTransaction(%v, %v, %v, %v, %v)", tt.args.transactionCode, tt.args.currency, tt.args.quantity, tt.args.transmitter, tt.args.transactionDate)
		})
	}
}

func Test_transactionService_ReplaceTransaction(t1 *testing.T) {
	transactions := []domain.Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
	}

	fakeTransaction := domain.Transaction{
		TransactionCode: "TC",
		Currency:        "COP",
		Quantity:        500,
		Transmitter:     "Central bank",
		TransactionDate: time.Time{},
	}

	funStore := mocks.GetFunStoreStub(transactions)
	errorStore := mocks.GetStoreStubWithErrors()

	funRepository := GetTransactionRepository(funStore)
	errorRepository := GetTransactionRepository(errorStore)

	type fields struct {
		transactionRepository TransactionRepository
	}
	type args struct {
		id          int
		code        string
		currency    string
		quantity    float64
		transmitter string
		date        time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Transaction
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Fun case for replace",
			fields{funRepository},
			args{1, fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{Id: 1, TransactionCode: fakeTransaction.TransactionCode, Currency: fakeTransaction.Currency, Quantity: fakeTransaction.Quantity, Transmitter: fakeTransaction.Transmitter, TransactionDate: fakeTransaction.TransactionDate},
			assert.NoError,
		},
		{
			"Error in store case for replace",
			fields{errorRepository},
			args{1, fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{},
			assert.Error,
		},
		{
			"Error case when id doesn't exist for replace",
			fields{funRepository},
			args{100, fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{},
			assert.Error,
		},
		{
			"Error case for replace transaction (error in repository)",
			fields{GetRepositoryMock(funStore, RepositoryMockArgs{HaveErrorInReplace: true})},
			args{1, fakeTransaction.TransactionCode, fakeTransaction.Currency, fakeTransaction.Quantity, fakeTransaction.Transmitter, fakeTransaction.TransactionDate},
			domain.Transaction{},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := transactionService{
				transactionRepository: tt.fields.transactionRepository,
			}
			got, err := t.ReplaceTransaction(tt.args.id, tt.args.code, tt.args.currency, tt.args.quantity, tt.args.transmitter, tt.args.date)
			if !tt.wantErr(t1, err, fmt.Sprintf("ReplaceTransaction(%v, %v, %v, %v, %v, %v)", tt.args.id, tt.args.code, tt.args.currency, tt.args.quantity, tt.args.transmitter, tt.args.date)) {
				return
			}
			assert.Equalf(t1, tt.want, got, "ReplaceTransaction(%v, %v, %v, %v, %v, %v)", tt.args.id, tt.args.code, tt.args.currency, tt.args.quantity, tt.args.transmitter, tt.args.date)
		})
	}
}
