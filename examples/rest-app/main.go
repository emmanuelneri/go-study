package main

import (
	"app/handler"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const (
	httpport   = ":8080"
	dbhost     = "localhost"
	dbport     = 5432
	dbuser     = "postgres"
	dbpassword = "postgres"
	dbname     = "goapp"
)

func main() {
	fmt.Println("App Go started", time.Now().Format("02-01-2006"))

	db := getDB()
	orderHandler := &handler.OrderHandler{Db: db}

	http.HandleFunc("/orders", orderHandler.Handle)
	http.ListenAndServe(httpport, nil)
}

func getDB() *sql.DB {
	dbConfig := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbuser, dbpassword, dbname)

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	return db
}
