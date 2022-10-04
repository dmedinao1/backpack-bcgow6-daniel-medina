package main

import (
	"github.com/dmedinao1/ejercicio-TM-04-10/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	transactionsGroup := app.Group("/transacciones")

	routes.DefineTransactionsRoutes(transactionsGroup)

	err := app.Run()

	if err != nil {
		panic(err)
	}
}
