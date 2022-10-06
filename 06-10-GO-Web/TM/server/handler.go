package server

import (
	"errors"
	"github.com/dmedinao1/go-web-practica/internal"
	"github.com/dmedinao1/go-web-practica/pkg/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
		if writeTokenErrorIfInvalid(c) {
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil || id < 1 {
			c.JSON(http.StatusBadRequest, web.New(http.StatusBadRequest, nil, errors.New("id inválido")))
			return
		}

		var transactionToSet fullTransactionRequest

		if err := c.ShouldBind(&transactionToSet); err != nil {
			c.JSON(http.StatusBadRequest, web.New(http.StatusBadRequest, nil, err))
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
			handleServiceError(c, err)
			return
		}

		c.JSON(http.StatusOK, web.New(http.StatusOK, newTransaction, nil))
	}
}

func (t transactionHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		if writeTokenErrorIfInvalid(c) {
			return
		}

		transactions, err := t.service.FindAll()

		if err != nil {
			handleServiceError(c, err)
			return
		}

		c.JSON(http.StatusOK, web.New(http.StatusOK, transactions, nil))
	}
}

func (t transactionHandler) SaveTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if writeTokenErrorIfInvalid(c) {
			return
		}

		var currentRequest fullTransactionRequest

		if err := c.ShouldBind(&currentRequest); err != nil {
			c.JSON(http.StatusBadRequest, web.New(http.StatusBadRequest, nil, err))

			return
		}

		savedTransaction, err := t.service.SaveTransaction(
			currentRequest.TransactionCode,
			currentRequest.Currency,
			currentRequest.Quantity,
			currentRequest.Transmitter,
			currentRequest.TransactionDate,
		)

		if err != nil {
			handleServiceError(c, err)
			return
		}

		c.JSON(http.StatusOK, web.New(http.StatusOK, savedTransaction, nil))
	}
}

func (t transactionHandler) UpdateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if writeTokenErrorIfInvalid(c) {
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil || id < 1 {
			c.JSON(http.StatusBadRequest, web.New(http.StatusBadRequest, nil, errors.New("id inválido")))
			return
		}

		var request codeAndQuantityRequest

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, web.New(http.StatusBadRequest, nil, err))
			return
		}

		err = t.service.UpdateCodeAndQuantityById(id, request.TransactionCode, request.Quantity)

		if err != nil {
			handleServiceError(c, err)
			return
		}

		c.JSON(http.StatusOK, web.New(http.StatusOK, request, nil))
	}
}

func (t transactionHandler) DeleteTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		if writeTokenErrorIfInvalid(c) {
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil || id < 1 {
			c.JSON(http.StatusBadRequest, web.New(http.StatusBadRequest, nil, errors.New("id inválido")))
			return
		}

		err = t.service.DeleteById(id)

		if err != nil {
			handleServiceError(c, err)
			return
		}

		c.JSON(http.StatusOK, web.New(http.StatusOK, gin.H{"id": id}, nil))
	}
}

func handleServiceError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case internal.ApiError:
		apiError := err.(internal.ApiError)
		statusCode = apiError.Code
	default:
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})

	return
}

func writeTokenErrorIfInvalid(c *gin.Context) bool {
	if !hasAValidToken(c) {
		c.JSON(http.StatusUnauthorized, web.New(http.StatusUnauthorized, nil, errors.New("token inválido")))

		return true
	}
	return false
}

func hasAValidToken(c *gin.Context) bool {
	token := c.GetHeader("token")
	tokenEnv := os.Getenv("TOKEN")
	return token == tokenEnv
}
