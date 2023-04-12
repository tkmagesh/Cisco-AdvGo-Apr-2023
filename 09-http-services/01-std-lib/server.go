package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products []Product = []Product{
	{Id: 1, Name: "Pen", Cost: 10},
	{Id: 2, Name: "Pencil", Cost: 5},
	{Id: 3, Name: "Marker", Cost: 50},
}

type appServer struct {
}

func (s *appServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		// fmt.Fprintln(w, "All the products will be served")
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			newProduct.Id = len(products) + 1
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "All the customers will be served")
	default:
		http.Error(w, "Resource not found", http.StatusNotFound)
	}

}

func main() {
	svr := &appServer{}
	http.ListenAndServe(":8080", svr)
}
