package main

import (
	"app/config"
	"app/internal/http"
	"context"
)

const (
	DBHost       = "localhost"
	DBPort       = "5432"
	DBUser       = "postgres"
	DBPassword   = "postgres"
	DBName       = "goapp"
	DBDriver     = "postgres"
	maxOpenConns = 10
)

func main() {
	db := config.StartDB(config.DBConfig{
		Host:         DBHost,
		Port:         DBPort,
		Name:         DBName,
		User:         DBUser,
		Password:     DBPassword,
		Driver:       DBDriver,
		MaxOpenConns: maxOpenConns,
		MigrationDir: "file://config/migrations",
	})
	ctx := context.WithValue(context.Background(), config.DB, db)
	http.Start(ctx)
}
