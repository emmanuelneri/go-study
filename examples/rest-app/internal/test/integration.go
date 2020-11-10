package test

import (
	"app/config"
	"context"
	"database/sql"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	postgresDockerImage = "postgres:12-alpine"
)

type Integration struct {
	context           context.Context
	postgresContainer testcontainers.Container
	dbConfig          config.DBConfig
	db                *sql.DB
}

func NewIntegration(dbConfig config.DBConfig) *Integration {
	return &Integration{dbConfig: dbConfig}
}

func (it *Integration) StartPostgres() error {
	it.context = context.Background()

	container, err := testcontainers.GenericContainer(it.context, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:      postgresDockerImage,
			Entrypoint: nil,
			Env: map[string]string{
				"POSTGRES_DB":       it.dbConfig.Name,
				"POSTGRES_USER":     it.dbConfig.User,
				"POSTGRES_PASSWORD": it.dbConfig.Password,
			},
			ExposedPorts: []string{"5432/tcp"},
			WaitingFor: wait.ForSQL("5432/tcp", "postgres", func(port nat.Port) string {
				it.dbConfig.Port = port.Port()
				return it.dbConfig.Url()
			}),
		},
		Started: true,
	})

	if err != nil {
		return err
	}

	it.postgresContainer = container
	return err
}

func (it *Integration) Terminate() error {
	return it.postgresContainer.Terminate(it.context)
}

func (it *Integration) StartDB() error {
	db, err := config.Connect(it.dbConfig)
	if err != nil {
		return err
	}

	it.db = db
	return nil
}

func (it Integration) DB() *sql.DB {
	return it.db
}
