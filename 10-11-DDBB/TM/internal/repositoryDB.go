package internal

import (
	"database/sql"
	"github.com/dmedinao1/go-web-practica/internal/domain"
)

const (
	FIND_ALL     = "SELECT id, transaction_code, currency, quantity, transmitter, transaction_date FROM transaction"
	INSERT       = "INSERT INTO transaction(id, transaction_code, currency, quantity, transmitter, transaction_date) VALUES (?, ?, ?, ?, ?, ?)"
	MAX_ID       = "SELECT MAX(id) FROM transaction"
	FIND_BY_ID   = "SELECT id, transaction_code, currency, quantity, transmitter, transaction_date FROM transaction WHERE id = ?"
	UPDATE       = "UPDATE transaction SET transaction_code = ?, currency = ?, quantity = ?, transmitter = ?, transaction_date = ? WHERE id = ?"
	DELETE_BY_ID = "DELETE FROM transaction WHERE id = ?"
)

func GetTransactionDBRepository(db *sql.DB) TransactionRepository {
	return &repository{db}
}

type repository struct {
	DB *sql.DB
}

func (r *repository) FindAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction

	rows, err := r.DB.Query(FIND_ALL)

	if err != nil {
		return transactions, err
	}

	for rows.Next() {
		var transaction domain.Transaction
		err := scanTransaction(&transaction, rows)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *repository) Save(transaction domain.Transaction) (domain.Transaction, error) {
	statement, err := r.DB.Prepare(INSERT)

	if err != nil {
		return domain.Transaction{}, err
	}

	defer statement.Close()

	_, err = statement.Exec(transaction.Id, transaction.TransactionCode, transaction.Currency, transaction.Quantity, transaction.Transmitter, transaction.TransactionDate)

	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (r *repository) GetLastId() (int, error) {
	var lastId int

	result, err := r.DB.Query(MAX_ID)

	if err != nil {
		return 0, err
	}

	err = result.Scan(&lastId)

	if err != nil {
		return 0, nil
	}

	return lastId, nil
}

func (r *repository) GetById(id int) (*domain.Transaction, error) {
	var transaction domain.Transaction

	rows, err := r.DB.Query(FIND_BY_ID, id)

	if err != nil {
		return &transaction, err
	}

	for rows.Next() {
		var transaction domain.Transaction
		err := scanTransaction(&transaction, rows)

		if err != nil {
			return nil, err
		}

		return &transaction, nil
	}

	return &transaction, nil
}

func (r *repository) Replace(id int, transaction domain.Transaction) (domain.Transaction, error) {
	statement, err := r.DB.Prepare(UPDATE)

	if err != nil {
		return domain.Transaction{}, err
	}

	defer statement.Close()

	_, err = statement.Exec(transaction.TransactionCode, transaction.Currency, transaction.Quantity, transaction.Transmitter, transaction.TransactionDate, id)

	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (r *repository) DeleteById(id int) error {
	stmt, err := r.DB.Prepare(DELETE_BY_ID)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func scanTransaction(toScan *domain.Transaction, rows *sql.Rows) error {
	return rows.Scan(&toScan.Id, &toScan.TransactionCode, &toScan.Currency, &toScan.Quantity, &toScan.Transmitter, &toScan.TransactionDate)
}
