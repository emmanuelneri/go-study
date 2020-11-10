package main

import (
	"app/config"
	"app/internal/container"
	"app/internal/http"
)

const (
	dbHost       = "localhost"
	dbPort       = "5432"
	dbUser       = "postgres"
	dbPassword   = "postgres"
	dbName       = "goapp"
	dbDriver     = "postgres"
	maxOpenConns = 10
	serverPort   = ":8080"
)

func main() {
	db := config.StartDB(config.DBConfig{
		Host:         dbHost,
		Port:         dbPort,
		Name:         dbName,
		User:         dbUser,
		Password:     dbPassword,
		Driver:       dbDriver,
		MaxOpenConns: maxOpenConns,
		MigrationDir: "file://config/migrations",
	})

	c := container.Start(db)

	http.NewHttpServer(serverPort, c).Start()
}
