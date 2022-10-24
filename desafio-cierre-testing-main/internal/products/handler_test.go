package products

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetFunStubService() Service {
	return &funService{}
}

type funService struct {
}

func (f *funService) GetAllBySeller(sellerID string) ([]Product, error) {
	return []Product{}, nil
}

func GetErrorStubService() Service {
	return &errorService{}
}

type errorService struct {
}

func (f *errorService) GetAllBySeller(sellerID string) ([]Product, error) {
	return nil, errors.New("service error")
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func createServer(service Service) *gin.Engine {
	handler := NewHandler(service)

	r := gin.Default()

	r.GET("", handler.GetProducts)
	return r
}

func Test_GetAll_OK(t *testing.T) {
	repo := NewRepository()
	service := NewService(repo)
	r := createServer(service)
	req, rr := createRequest(http.MethodGet, "/?seller_id=FEX112AC", ``)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_GetAll_BadRequest(t *testing.T) {
	r := createServer(GetFunStubService())
	req, rr := createRequest(http.MethodGet, "/", ``)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_GetAll_InternalError(t *testing.T) {
	r := createServer(GetErrorStubService())
	req, rr := createRequest(http.MethodGet, "/?seller_id=FEX112AC", ``)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
