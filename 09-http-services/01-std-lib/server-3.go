// To be fixed
package main

import (
	"encoding/json"
	"fmt"
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

// type HandlerFn func(http.ResponseWriter, *http.Request)

type appRouter struct {
	routes map[string]http.HandlerFunc
}

func (router *appRouter) Register(pattern string, handlerFn http.HandlerFunc) {
	router.routes[pattern] = handlerFn
}

func (router *appRouter) Method(methodName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == methodName {
			router.ServeHTTP(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (router *appRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resourePath := r.URL.Path
	handlerFn := router.routes[resourePath]
	if handlerFn != nil {
		handlerFn(w, r)
		return
	}
	http.Error(w, "Resource not found", http.StatusNotFound)
}

func NewAppRouter() *appRouter {
	return &appRouter{
		routes: make(map[string]http.HandlerFunc),
	}
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
	router := NewAppRouter()
	router.Register("/", indexGetHandler).Method(http.MethodGet)
	router.Register("/products", productsHandler)
	router.Register("/customers", customersHandler)
	router.Register("/users", usersHandler)
	http.ListenAndServe(":8080", router)
}
