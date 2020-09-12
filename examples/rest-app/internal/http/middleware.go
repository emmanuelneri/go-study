package http

import (
	"log"
	"net/http"
	"time"
)

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path %s requested", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func measureRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Path %s took %v", r.RequestURI, time.Since(start))
	})
}
