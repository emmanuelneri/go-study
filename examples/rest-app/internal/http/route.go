package http

import (
	"app/internal/handler"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

const PostMethod = "POST"

func configRoutes(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	orderHandler := handler.NewOrderHandler(ctx)
	r.Handle(
		"/orders",
		addDefaultMiddlewares(orderHandler.Handle),
	).Methods(PostMethod)

	return r
}

func addDefaultMiddlewares(httpFunc func(http.ResponseWriter, *http.Request)) http.Handler {
	return measureRequest(logRequest(http.HandlerFunc(httpFunc)))
}
