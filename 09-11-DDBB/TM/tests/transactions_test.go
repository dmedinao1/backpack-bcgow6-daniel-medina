package tests

import (
	"bytes"
	"github.com/dmedinao1/go-web-practica/internal"
	"github.com/dmedinao1/go-web-practica/pkg/store"
	"github.com/dmedinao1/go-web-practica/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func createServer() *gin.Engine {
	_ = godotenv.Load()
	appStore := store.New(store.FileType, "/Users/danmedina/Documents/github/backpack-bcgow6-daniel-medina/myFiles/data.json")
	transactionRepository := internal.GetTransactionRepository(appStore)
	transactionService := internal.GetTransactionService(transactionRepository)
	transactionHandlers := server.GetTransactionHandler(transactionService)
	r := gin.Default()

	pr := r.Group("/transactions")
	pr.PATCH("/:id", transactionHandlers.UpdateTransaction())
	pr.DELETE("/:id", transactionHandlers.DeleteTransaction())
	return r
}

func Test_Update_OK(t *testing.T) {
	r := createServer()
	req, rr := createRequest(http.MethodPatch, "/transactions/1", `{"transactionCode": "TC-100","quantity": 100200.1}`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
func Test_Delete_OK(t *testing.T) {
	r := createServer()
	req, rr := createRequest(http.MethodDelete, "/transactions/3", ``)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
