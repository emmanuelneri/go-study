package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const (
	address = ":8080"
)

func Start(ctx context.Context) {
	fmt.Println("Starting HTTP Server at " + address)

	r := configRoutes(ctx)
	http.Handle("/", r)

	log.Panicln(http.ListenAndServe(address, nil))
}
