package main

import (
	"app/config"
	internalHttp "app/internal/http"
	"app/internal/test"
	"app/pkg/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gotest.tools/assert"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

const (
	url = "http://localhost:8080/orders"
)

func TestApp(t *testing.T) {
	dbConfig := config.DBConfig{
		Host:         DBHost,
		Port:         DBPort,
		Name:         DBName,
		User:         DBUser,
		Password:     DBPassword,
		Driver:       DBDriver,
		MaxOpenConns: maxOpenConns,
		MigrationDir: "file://../config/migrations",
	}

	testContext := test.NewIntegration(dbConfig)
	defer testContext.Terminate()

	if err := testContext.StartPostgres(); err != nil {
		t.Error(err)
	}

	if err := testContext.StartAppContext(); err != nil {
		t.Error(err)
	}

	go internalHttp.Start(testContext.AppContext())

	request := model.Order{
		ID:       rand.Int(),
		Customer: fmt.Sprintf("Customer %d", rand.Int()),
		Total:    rand.Float32(),
	}

	response, err := send(request)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, response.StatusCode)
}

func send(order model.Order) (*http.Response, error) {
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
