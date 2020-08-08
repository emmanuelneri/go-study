package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	model "app/model"
)

// Order type to represent model
type Order model.Order

// OrderHandler use to get db address for HTTP Handle method set
type OrderHandler struct {
	Db *sql.DB
}

// Handle represent Order HTTP Post request
func (orderHandler *OrderHandler) Handle(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(responseWriter, "404 not found.", http.StatusNotFound)
		return
	}

	if request.Body == nil {
		http.Error(responseWriter, "body required.", http.StatusBadRequest)
	}

	var order Order
	err := json.NewDecoder(request.Body).Decode(&order)

	if err != nil {
		log.Panicln("fail to process order ", err)
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	defer request.Body.Close()
	log.Print("Order received: ", order)
	saveError := save(&order, orderHandler)

	if saveError != nil {
		log.Panicln("fail to save order ", err)
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	log.Print("Order saved: ", order)
	json.NewEncoder(responseWriter).Encode(order)
}
func save(order *Order, orderHandler *OrderHandler) error {
	var id int
	rowError := orderHandler.Db.QueryRow("INSERT INTO sales_order(customer, total) VALUES($1,$2) RETURNING id", order.Customer, order.Total).Scan(&id)
	order.ID = id
	return rowError
}
