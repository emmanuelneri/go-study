package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Http start")

	// http://localhost:8090/queryParam?Param=Test
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(responseWriter, "Hello HTTP GET")
	})

	// http://localhost:8090/queryParam?Param=Test
	http.HandleFunc("/queryParam", func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Println("Query request: ", request.URL.Query())
		fmt.Println("Param Value: ", request.URL.Query()["Param"])
	})

	// http://localhost:8090/pathParam/1
	http.HandleFunc("/pathParam/", func(responseWriter http.ResponseWriter, request *http.Request) {
		id := strings.TrimPrefix(request.URL.Path, "/pathParam/")
		fmt.Println("Id requested: ", id)
	})

	http.ListenAndServe(":8090", nil)
}
