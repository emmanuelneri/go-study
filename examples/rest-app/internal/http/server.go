package http

import (
	"app/internal/container"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type HttpServer struct {
	port      string
	container *container.Container
}

func NewHttpServer(port string, container *container.Container) *HttpServer {
	return &HttpServer{
		port:      port,
		container: container,
	}
}

func (h HttpServer) Start() {
	fmt.Println("Starting HTTP Server at " + h.port)

	r := configRoutes(h.container)
	http.Handle("/", r)

	log.Panicln(http.ListenAndServe(h.port, nil))
}
