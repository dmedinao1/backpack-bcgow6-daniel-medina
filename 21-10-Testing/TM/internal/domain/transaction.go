package domain

import "time"

type Transaction struct {
	Id              int       `json:"id"`
	TransactionCode string    `json:"transactionCode"`
	Currency        string    `json:"currency"`
	Quantity        float64   `json:"quantity"`
	Transmitter     string    `json:"transmitter"`
	TransactionDate time.Time `json:"transactionDate"`
}
