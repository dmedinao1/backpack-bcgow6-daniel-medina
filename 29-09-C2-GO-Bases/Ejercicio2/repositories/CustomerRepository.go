package repositories

import (
	"fmt"
	"os"
	"strings"

	"github.com/dmedinao1/backpack-bcgow6-daniel-medina/customers/models"
)

func getFilePath() string {
	workingDirectory, _ := os.Getwd()
	return fmt.Sprintf("%s/%s/%s", workingDirectory, "myFiles", "customers.txt")
}

func getRawCustomers() []string {
	data, err := os.ReadFile(getFilePath())

	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func GetAllRegisteredCustomers() *[]models.Customer {
	var customers []models.Customer
	rawCustomers := getRawCustomers()

	for _, rawCustomer := range rawCustomers {
		customer := models.Customer{}
		customer.FillFromString(rawCustomer)
		customers = append(customers, customer)
	}

	return &customers
}

func GenerateNextId() int {
	registeredCostumers := GetAllRegisteredCustomers()
	var currentId int

	if len(*registeredCostumers) > 0 {
		currentId = (*registeredCostumers)[0].Id
	}

	return currentId + 1
}

func SaveCustomer(c *models.Customer) {
	saved := c.ToString()
	rawCustomers := strings.Join(getRawCustomers(), "\n")
	rawCustomers += saved
	err := os.WriteFile(getFilePath(), []byte(rawCustomers), 0644)

	if err != nil {
		panic(err)
	}
}
