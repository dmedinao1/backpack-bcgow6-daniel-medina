package main

import (
	"github.com/dmedinao1/ejercicio-TT-04-10/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	transactionRepository := internal.GetTransactionRepository()
	transactionService := internal.GetTransactionService(transactionRepository)
	transactionHandlers := GetTransactionHandler(transactionService)

	app := gin.Default()

	transactionsRoutes := app.Group("/transacciones")

	transactionsRoutes.POST("", transactionHandlers.SaveTransaction())
	transactionsRoutes.GET("", transactionHandlers.GetAll())

	err := app.Run()

	if err != nil {
		panic(err)
	}

}
