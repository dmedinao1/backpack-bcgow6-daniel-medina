package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	rawData, err := os.ReadFile(jsonPath)

	if err != nil {
		fmt.Println("No se pudo leer el archivo", err)
		return
	}

	var transactions []transaction

	err = json.Unmarshal(rawData, &transactions)

	if err != nil {
		fmt.Println("No se pudo convertir el archivo", err)
		return
	}

	fmt.Printf("%+v\n", transactions)

}
