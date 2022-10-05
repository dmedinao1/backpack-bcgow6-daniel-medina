package server

import (
	"fmt"
	"github.com/dmedinao1/go-web-practica/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type fullTransactionRequest struct {
	TransactionCode string    `json:"transactionCode" binding:"required"`
	Currency        string    `json:"currency" binding:"required"`
	Quantity        float64   `json:"quantity" binding:"required"`
	Transmitter     string    `json:"transmitter" binding:"required"`
	TransactionDate time.Time `json:"transactionDate" binding:"required"`
}

type codeAndQuantityRequest struct {
	TransactionCode string  `json:"transactionCode" binding:"required"`
	Quantity        float64 `json:"quantity" binding:"required"`
}

func GetTransactionHandler(service internal.TransactionService) TransactionHandlers {
	return transactionHandler{service: service}
}

type TransactionHandlers interface {
	GetAll() gin.HandlerFunc
	SaveTransaction() gin.HandlerFunc
	ReplaceTransaction() gin.HandlerFunc
	UpdateTransaction() gin.HandlerFunc
	DeleteTransaction() gin.HandlerFunc
}

type transactionHandler struct {
	service internal.TransactionService
}

func (t transactionHandler) ReplaceTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !hasAValidToken(c) {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})

			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Id inválido"),
			})

			return
		}

		var transactionToSet fullTransactionRequest

		if err := c.ShouldBind(&transactionToSet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		newTransaction, err := t.service.ReplaceTransaction(
			id,
			transactionToSet.TransactionCode,
			transactionToSet.Currency,
			transactionToSet.Quantity,
			transactionToSet.Transmitter,
			transactionToSet.TransactionDate,
		)

		if err != nil {
			var statusCode int

			switch err.(type) {
			case *internal.ApiError:
				apiError := err.(*internal.ApiError)
				statusCode = apiError.Code
			default:
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, newTransaction)
	}
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

		var currentRequest fullTransactionRequest

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

func (t transactionHandler) UpdateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !hasAValidToken(c) {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})

			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Id inválido"),
			})

			return
		}

		var request codeAndQuantityRequest

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = t.service.UpdateCodeAndQuantityById(id, request.TransactionCode, request.Quantity)

		if err != nil {
			var statusCode int

			switch err.(type) {
			case *internal.ApiError:
				apiError := err.(*internal.ApiError)
				statusCode = apiError.Code
			default:
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, request)
	}
}

func (t transactionHandler) DeleteTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !hasAValidToken(c) {
			c.JSON(401, gin.H{
				"error": "Invalid token",
			})

			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil || id < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Id inválido"),
			})

			return
		}

		err = t.service.DeleteById(id)

		if err != nil {
			var statusCode int

			switch err.(type) {
			case *internal.ApiError:
				apiError := err.(*internal.ApiError)
				statusCode = apiError.Code
			default:
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

func hasAValidToken(c *gin.Context) bool {
	token := c.GetHeader("token")
	return token == "1234"
}
