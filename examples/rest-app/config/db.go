package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	dbUser       = "postgres"
	dbPassword   = "postgres"
	dbName       = "goapp"
	driver       = "postgres"
	maxOpenConns = 10
)

func StartDB() *sql.DB {
	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	fmt.Println("DB starting... " + dbConfig)

	db, err := sql.Open(driver, dbConfig)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenConns)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
