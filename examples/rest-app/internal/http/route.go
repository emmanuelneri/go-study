package http

import (
	"app/internal/handler"
	"context"
	"github.com/gorilla/mux"
)

const PostMethod = "POST"

func configRoutes(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	orderHandler := handler.NewOrderHandler(ctx)
	r.HandleFunc("/orders", orderHandler.Handle).Methods(PostMethod)

	return r
}
