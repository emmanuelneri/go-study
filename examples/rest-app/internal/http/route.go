package http

import (
	"app/internal/container"
	"app/internal/handler"
	"github.com/gorilla/mux"
	"net/http"
)

const PostMethod = "POST"

func configRoutes(container *container.Container) *mux.Router {
	r := mux.NewRouter()
	orderHandler := handler.NewOrderHandler(container.OrderService)
	r.Handle(
		"/orders",
		addDefaultMiddlewares(orderHandler.Handle),
	).Methods(PostMethod)

	return r
}

func addDefaultMiddlewares(httpFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return measureRequest(logRequest(http.HandlerFunc(httpFunc)))
}
