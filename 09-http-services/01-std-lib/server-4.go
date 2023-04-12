// To be fixed
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

type Middleware func(http.HandlerFunc) http.HandlerFunc

func logHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func profileHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func chainHandlers(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World! - [GET]")
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World! - [POST]")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers will be served")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the users will be served")
}

func main() {
	// svr := &appServer{}
	/*
		router := http.DefaultServeMux
		router.HandleFunc("/", indexGetHandler)
		router.HandleFunc("/products", productsHandler)
		router.HandleFunc("/customers", customersHandler)
		router.HandleFunc("/users", usersHandler)
		http.ListenAndServe(":8080", nil)
	*/

	/*
		http.HandleFunc("/", profileHandler(logHandler(indexGetHandler)))
		http.HandleFunc("/products", profileHandler(logHandler(productsHandler)))
		http.HandleFunc("/customers", profileHandler(logHandler(customersHandler)))
		http.HandleFunc("/users", profileHandler(logHandler(usersHandler))) */

	http.HandleFunc("/", chainHandlers(indexGetHandler, logHandler, profileHandler))
	http.HandleFunc("/products", chainHandlers(productsHandler, logHandler, profileHandler))
	http.HandleFunc("/customers", chainHandlers(customersHandler, logHandler, profileHandler))
	http.HandleFunc("/users", chainHandlers(usersHandler, logHandler, profileHandler))
	http.ListenAndServe(":8080", nil)
}
