package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Http start")

	// http://localhost:8090/hello
	http.HandleFunc("/hello", hello())

	// http://localhost:8090/queryParam?Param=Test
	http.HandleFunc("/queryParam", queryParam())

	// http://localhost:8090/pathParam/1
	http.HandleFunc("/pathParam/", pathParam())

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("server failed. ", err)
	}
}

func hello() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprint(responseWriter, "Hello HTTP")
	}
}

func queryParam() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Println("Query request: ", request.URL.Query())
		fmt.Println("Param Value: ", request.URL.Query()["Param"])
	}
}

func pathParam() func(responseWriter http.ResponseWriter, request *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		id := strings.TrimPrefix(request.URL.Path, "/pathParam/")
		fmt.Println("Id requested: ", id)
	}
}
