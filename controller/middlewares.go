package controller

import (
	"log"
	"net/http"
)

func logRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %v to %v", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		w.Header()
		log.Printf("Respond %v to %v", r.Method, r.RequestURI)
	})
}
