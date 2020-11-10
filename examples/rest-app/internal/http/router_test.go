package http

import (
	"app/config"
	"app/internal/container"
	"app/internal/test"
	"app/pkg/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

const (
	serverPort   = ":9090"
	url          = "http://localhost:9090/orders"
	dbHost       = "localhost"
	dbPort       = "5432"
	dbUser       = "postgres"
	dbPassword   = "postgres"
	dbName       = "goapp"
	dbDriver     = "postgres"
	maxOpenConns = 10
)

func TestApp(t *testing.T) {
	dbConfig := config.DBConfig{
		Host:         dbHost,
		Port:         dbPort,
		Name:         dbName,
		User:         dbUser,
		Password:     dbPassword,
		Driver:       dbDriver,
		MaxOpenConns: maxOpenConns,
		MigrationDir: "file://../../config/migrations",
	}

	testContext := test.NewIntegration(dbConfig)
	defer testContext.Terminate()

	if err := testContext.StartPostgres(); err != nil {
		t.Error(err)
	}

	if err := testContext.StartDB(); err != nil {
		t.Error(err)
	}

	c := container.Start(testContext.DB())
	go func() {
		r := configRoutes(c)
		http.Handle("/", r)
		log.Panicln(http.ListenAndServe(serverPort, nil))
	}()

	request := domain.Order{
		Customer: fmt.Sprintf("Customer %d", rand.Int()),
		Total:    19.99,
	}

	response, err := send(request)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
	orderResponse := domain.Order{}
	err = json.NewDecoder(response.Body).Decode(&orderResponse)
	assert.Nil(t, err)

	assert.Equal(t, request.Customer, orderResponse.Customer)
	assert.Equal(t, request.Total, orderResponse.Total)

	var customerSaved string
	var totalSaved float32
	err = testContext.DB().QueryRow("SELECT customer, total FROM sales_order WHERE id = $1", orderResponse.ID).Scan(&customerSaved, &totalSaved)
	assert.Nil(t, err)

	assert.Equal(t, orderResponse.Customer, customerSaved)
	assert.Equal(t, orderResponse.Total, totalSaved)
}

func send(order domain.Order) (*http.Response, error) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	body, err := json.Marshal(order)
	if err != nil {
		return nil, errors.Wrap(err, "marshal error")
	}

	return client.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
}
