package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"./structs"
)

type Item structs.Item
type Order structs.Order

func main() {
	const startMessage = "App Go started"

	fmt.Println(startMessage, time.Now().Format("02-01-2006"))
	fmt.Println("------------")

	http.HandleFunc("/orders", handleOrder)
	http.ListenAndServe(":8080", nil)
}

func handleOrder(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(responseWriter, "404 not found.", http.StatusNotFound)
		return
	}

	var order Order
	err := json.NewDecoder(request.Body).Decode(&order)

	if err != nil {
		log.Panicln("fail to process order ", err)
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	log.Print("Order received: ", order)
}
