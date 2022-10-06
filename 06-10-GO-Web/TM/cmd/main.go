package main

import (
	"github.com/dmedinao1/go-web-practica/internal"
	"github.com/dmedinao1/go-web-practica/pkg/store"
	"github.com/dmedinao1/go-web-practica/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()
	appStore := store.New(store.FileType, os.Getenv("JSON_STORE_FILE"))
	transactionRepository := internal.GetTransactionRepository(appStore)
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
