package main

import (
	"github.com/dmedinao1/go-web-practica/docs"
	"github.com/dmedinao1/go-web-practica/internal"
	"github.com/dmedinao1/go-web-practica/pkg/db"
	"github.com/dmedinao1/go-web-practica/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mercado libre bootcamp | Transacciones API
// @version 1.0
// @description Servicios para crear, listar, actualizar y eliminar transacciones.

// @contact.name   Daniel Medina
// @contact.url    http://github.com/dmedinao1
// @contact.email  daniel.medina@mercadolibre.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()

	dbConnection, err := db.GetDBConnection()

	if err != nil {
		panic(err)
	}

	transactionRepository := internal.GetTransactionDBRepository(dbConnection)
	transactionService := internal.GetTransactionService(transactionRepository)
	transactionHandlers := server.GetTransactionHandler(transactionService)

	app := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	transactionsRoutes := app.Group("/transacciones")

	transactionsRoutes.POST("", transactionHandlers.SaveTransaction())
	transactionsRoutes.GET("", transactionHandlers.GetAll())
	transactionsRoutes.PUT(":id", transactionHandlers.ReplaceTransaction())
	transactionsRoutes.PATCH(":id", transactionHandlers.UpdateTransaction())
	transactionsRoutes.DELETE(":id", transactionHandlers.DeleteTransaction())

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = app.Run()

	if err != nil {
		panic(err)
	}

}
