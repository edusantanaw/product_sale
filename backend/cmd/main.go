package main

import (
	"net/http"

	"github.com/edusantanaw/product_sale/internal"
)

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"message": "Hello, World!"}`))
	})
	serverMux.HandleFunc("POST /customers", internal.CreateCustomer)
	server := http.Server{
		Addr:    ":3000",
		Handler: serverMux,
	}
	server.ListenAndServe()
}
