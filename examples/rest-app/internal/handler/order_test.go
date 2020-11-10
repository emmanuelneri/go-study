package handler

import (
	"app/internal/container"
	"app/internal/repository/mocks"
	"app/internal/service"
	"app/pkg/domain"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestRequestWithEmptyBodyShouldBeReturnBadRequest(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)

	c := container.Start(db)
	orderHandler := NewOrderHandler(c.OrderService)

	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecoder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(orderHandler.Handle)
	testHandler.ServeHTTP(responseRecoder, req)

	assert.Equal(t, 400, responseRecoder.Code)
	assert.Equal(t, "body required\n", responseRecoder.Body.String())
}

func TestRequestWithInvalidBodyShouldBeReturnBadRequest(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)

	c := container.Start(db)
	orderHandler := NewOrderHandler(c.OrderService)

	body := strings.NewReader("invalid body")
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatal(err)
	}

	responseRecoder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(orderHandler.Handle)
	testHandler.ServeHTTP(responseRecoder, req)

	assert.Equal(t, 400, responseRecoder.Code)
}

func TestValidBodyShouldBeReturnOk(t *testing.T) {
	db, sqlMock, err := sqlmock.New()
	assert.Nil(t, err)

	id := 1
	requestBody := Request{
		Customer: fmt.Sprintf("Customer %d", rand.Int()),
		Total:    rand.Float32(),
	}

	sqlMock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO sales_order(customer, total) VALUES($1,$2) RETURNING id`)).
		WithArgs(requestBody.Customer, requestBody.Total).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

	c := container.Start(db)
	orderHandler := NewOrderHandler(c.OrderService)

	body, err := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	responseRecoder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(orderHandler.Handle)
	testHandler.ServeHTTP(responseRecoder, req)

	assert.Equal(t, 200, responseRecoder.Code)
	expected := fmt.Sprintf("{\"ID\":%d,\"Customer\":\"%s\",\"Total\":%v}\n", id, requestBody.Customer, requestBody.Total)
	assert.Equal(t, expected, responseRecoder.Body.String())
}

func TestSaveFailReturnShouldBeReturnBadRequest(t *testing.T) {
	requestBody := Request{
		Customer: fmt.Sprintf("Customer %d", rand.Int()),
		Total:    rand.Float32(),
	}

	order := &domain.Order{
		Customer: requestBody.Customer,
		Total:    requestBody.Total,
	}

	orderRepository := new(mocks.OrderRepository)
	orderRepository.On("Save", order).Return(errors.New("db unavailable"))
	orderService := service.OrderService(orderRepository)

	c := &container.Container{OrderService: orderService}
	orderHandler := NewOrderHandler(c.OrderService)

	body, err := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	responseRecoder := httptest.NewRecorder()
	testHandler := http.HandlerFunc(orderHandler.Handle)
	testHandler.ServeHTTP(responseRecoder, req)

	assert.Equal(t, 400, responseRecoder.Code)
}
