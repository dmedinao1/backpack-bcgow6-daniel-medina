package main

import (
	"github.com/dmedinao1/go-web-practica/internal"
	"github.com/dmedinao1/go-web-practica/server"
	"github.com/gin-gonic/gin"
)

func main() {
	transactionRepository := internal.GetTransactionRepository()
	transactionService := internal.GetTransactionService(transactionRepository)
	transactionHandlers := server.GetTransactionHandler(transactionService)

	app := gin.Default()

	transactionsRoutes := app.Group("/transacciones")

	transactionsRoutes.POST("", transactionHandlers.SaveTransaction())
	transactionsRoutes.GET("", transactionHandlers.GetAll())
	transactionsRoutes.PUT(":id", transactionHandlers.ReplaceTransaction())
	transactionsRoutes.PATCH(":id", transactionHandlers.UpdateTransaction())
	transactionsRoutes.DELETE(":id", transactionHandlers.DeleteTransaction())

	err := app.Run()

	if err != nil {
		panic(err)
	}

}
