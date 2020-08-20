package handler

import (
	"app/internal/repository"
	"context"
	"encoding/json"
	"log"
	"net/http"

	model "app/pkg/model"
)

type Order model.Order

type OrderHandler struct {
	Repository repository.OrderRepository
}

func NewOrderHandler(ctx context.Context) *OrderHandler {
	return &OrderHandler{Repository: repository.NewOrderRepositoryImpl(ctx)}
}

func (h *OrderHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(responseWriter, "body required.", http.StatusBadRequest)
	}

	order := &model.Order{}
	err := json.NewDecoder(request.Body).Decode(&order)

	if err != nil {
		log.Printf("fail to process order %s", err)
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	defer request.Body.Close()
	log.Printf("Order received: %v", order)
	orderRepository := h.Repository
	err = orderRepository.Save(order)

	if err != nil {
		log.Printf("fail to save order %s", err)
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Order saved: %v", order)
	_ = json.NewEncoder(responseWriter).Encode(order)
}
