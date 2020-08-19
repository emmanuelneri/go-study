package main

import (
	"app/config"
	"app/internal/http"
	"context"
)

func main() {
	db := config.StartDB()
	ctx := context.WithValue(context.Background(), config.DB, db)
	http.Start(ctx)
}
