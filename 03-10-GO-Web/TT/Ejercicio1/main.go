package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

const jsonPath = "/Users/danmedina/Documents/github/backpack-bcgow6-daniel-medina/03-10-GO-Web/TM/Ejercicio1/data.json"

type transaction struct {
	Id              int       `json:"id"`
	TransactionCode string    `json:"transactionCode"`
	Currency        string    `json:"currency"`
	Quantity        float64   `json:"quantity"`
	Transmitter     string    `json:"transmitter"`
	TransactionDate time.Time `json:"transactionDate"`
}

func (t transaction) getFromFile(filePath string) (*[]transaction, error) {
	var transactions []transaction

	rawData, err := os.ReadFile(filePath)

	if err != nil {
		return &transactions, err
	}

	err = json.Unmarshal(rawData, &transactions)

	if err != nil {
		return &transactions, err
	}

	return &transactions, nil
}

func (t transaction) checkFilters(params *transaction) bool {
	if params.Id != 0 && t.Id != params.Id {
		return false
	}

	if params.TransactionCode != "" && t.TransactionCode != params.TransactionCode {
		return false
	}

	if params.Currency != "" && t.Currency != params.Currency {
		return false
	}

	if params.Quantity != 0 && t.Quantity != params.Quantity {
		return false
	}

	if params.Transmitter != "" && t.Transmitter != params.Transmitter {
		return false
	}

	if !params.TransactionDate.IsZero() && t.TransactionDate.Equal(params.TransactionDate) {
		return false
	}

	return true
}

func main() {
	router := gin.Default()

	router.GET("/transacciones/busqueda", func(c *gin.Context) {

		searchParams, err := getSearchParams(c)

		if err != nil {
			c.JSON(400, gin.H{
				"error": fmt.Sprintf("no se pudo leer los parámetros error: %s", err.Error()),
			})
			return
		}

		transactions, err := transaction{}.getFromFile(jsonPath)

		var filteredTransactions []transaction

		for _, currentTransaction := range *transactions {
			ok := currentTransaction.checkFilters(searchParams)
			if ok {
				filteredTransactions = append(filteredTransactions, currentTransaction)
			}
		}

		if len(filteredTransactions) == 0 {
			c.JSON(404, gin.H{
				"error": "no se encontraron transacciones con los parámetros establecidos",
			})
			return
		}

		c.JSON(200, filteredTransactions)
	})

	router.Run()
}

func getSearchParams(c *gin.Context) (*transaction, error) {
	searchParams := &transaction{}

	id, okId := c.GetQuery("id")

	if okId {
		value, err := strconv.Atoi(id)

		if err != nil {
			return searchParams, err
		}

		searchParams.Id = value
	}

	code, okCode := c.GetQuery("code")

	if okCode {
		searchParams.TransactionCode = code
	}

	currency, okCurrency := c.GetQuery("currency")

	if okCurrency {
		searchParams.Currency = currency
	}

	quantity, okQuantity := c.GetQuery("quantity")

	if okQuantity {
		value, err := strconv.ParseFloat(quantity, 64)

		if err != nil {
			return searchParams, err
		}

		searchParams.Quantity = value
	}

	transmitter, okTransmitter := c.GetQuery("transmitter")

	if okTransmitter {
		searchParams.Transmitter = transmitter
	}

	date, okDate := c.GetQuery("date")

	if okDate {
		parsedDate, err := time.Parse(time.RFC3339, date)

		if err != nil {
			return nil, err
		}

		searchParams.TransactionDate = parsedDate
	}

	return searchParams, nil
}
