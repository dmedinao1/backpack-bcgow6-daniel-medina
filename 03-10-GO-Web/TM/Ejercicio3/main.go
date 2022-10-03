package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
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

func main() {
	router := gin.Default()

	router.GET("/transacciones", func(c *gin.Context) {

		transactions, err := transaction{}.getFromFile(jsonPath)

		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, transactions)
	})

	router.Run()
}
