package main

import (
	"github.com/dmedinao1/ejercicio-TT-04-10/internal"
	"github.com/gin-gonic/gin"
	"time"
)

type request struct {
	TransactionCode string    `json:"transactionCode" binding:"required"`
	Currency        string    `json:"currency" binding:"required"`
	Quantity        float64   `json:"quantity" binding:"required"`
	Transmitter     string    `json:"transmitter" binding:"required"`
	TransactionDate time.Time `json:"transactionDate" binding:"required"`
}

func GetTransactionHandler(service internal.TransactionService) TransactionHandlers {
	return transactionHandler{service: service}
}

type TransactionHandlers interface {
	GetAll() gin.HandlerFunc
	SaveTransaction() gin.HandlerFunc
}

type transactionHandler struct {
	service internal.TransactionService
}

func (t transactionHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !hasAValidToken(c) {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})

			return
		}

		c.JSON(200, t.service.FindAll())
	}
}

func (t transactionHandler) SaveTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !hasAValidToken(c) {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})

			return
		}

		var currentRequest request

		if err := c.ShouldBind(&currentRequest); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		savedTransaction := t.service.SaveTransaction(
			currentRequest.TransactionCode,
			currentRequest.Currency,
			currentRequest.Quantity,
			currentRequest.Transmitter,
			currentRequest.TransactionDate,
		)

		c.JSON(200, savedTransaction)
	}
}

func hasAValidToken(c *gin.Context) bool {
	token := c.GetHeader("token")
	return token == "1234"
}
