package handler

import (
	"app/internal/service"
	"app/pkg/domain"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Customer string
	Total    float32
}

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, errBodyRequired, http.StatusBadRequest)
		return
	}

	requestBody := &Request{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		log.Printf("[ERROR] fail to process order %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	order := &domain.Order{
		Customer: requestBody.Customer,
		Total:    requestBody.Total,
	}

	if err = h.service.Save(order); err != nil {
		log.Printf("[ERROR] fail to save order %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(order); err != nil {
		log.Printf("[ERROR] fail to encode response order %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Order saved: %v", order)
}
