package internal

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dmedinao1/go-web-practica/internal/domain"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func Test_repository_FindAll(t *testing.T) {
	// Arrange
	transactions := []domain.Transaction{
		{1, "TC1", "USD", 134.1, "Central Bank", time.Now()},
		{2, "TC2", "COP", 134.1, "Central Bank", time.Now()},
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "transaction_code", "currency", "quantity", "transmitter", "transaction_date"}
	rows := sqlmock.NewRows(columns)

	addTransactionsToRows(rows, transactions)

	mock.ExpectQuery(FIND_ALL).WillReturnRows(rows)

	dbError, mockError, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbError.Close()

	mockError.ExpectQuery(FIND_ALL).WillReturnError(assert.AnError)

	type fields struct {
		DB *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Transaction
		wantErr assert.ErrorAssertionFunc
	}{
		{"Fun case for get all", fields{db}, transactions, assert.NoError},
		{"Error case for get all", fields{dbError}, nil, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetTransactionDBRepository(tt.fields.DB)
			// Act
			got, err := r.FindAll()

			// Assert
			if !tt.wantErr(t, err, fmt.Sprintf("FindAll()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "FindAll()")
		})
	}
}

func Test_repository_Save(t *testing.T) {
	// Arrange
	transaction := domain.Transaction{Id: 1, TransactionCode: "TC1", Currency: "USD", Quantity: 134.1, Transmitter: "Central Bank", TransactionDate: time.Now()}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(INSERT))
	mock.ExpectExec(regexp.QuoteMeta(INSERT)).
		WithArgs(transaction.Id, transaction.TransactionCode, transaction.Currency, transaction.Quantity, transaction.Transmitter, transaction.TransactionDate).
		WillReturnResult(sqlmock.NewResult(int64(transaction.Id), 1))

	dbError, mockError, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbError.Close()

	mockError.ExpectPrepare(regexp.QuoteMeta(INSERT))
	mockError.ExpectExec(regexp.QuoteMeta(INSERT)).
		WithArgs(transaction.Id, transaction.TransactionCode, transaction.Currency, transaction.Quantity, transaction.Transmitter, transaction.TransactionDate).
		WillReturnError(assert.AnError)

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		transaction domain.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Transaction
		wantErr assert.ErrorAssertionFunc
	}{
		{"Fun case for save", fields{db}, args{transaction}, transaction, assert.NoError},
		{"Error case for save", fields{dbError}, args{transaction}, domain.Transaction{}, assert.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetTransactionDBRepository(tt.fields.DB)
			got, err := r.Save(tt.args.transaction)
			if !tt.wantErr(t, err, fmt.Sprintf("Save(%v)", tt.args.transaction)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Save(%v)", tt.args.transaction)
		})
	}
}

func addTransactionsToRows(rows *sqlmock.Rows, transactions []domain.Transaction) {
	for _, transaction := range transactions {
		rows.AddRow(transaction.Id, transaction.TransactionCode, transaction.Currency, transaction.Quantity, transaction.Transmitter, transaction.TransactionDate)
	}
}
