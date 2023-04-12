// To be fixed
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World! - [GET]")
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var product *Product
	if productId, err := strconv.Atoi(vars["product_id"]); err == nil {
		for _, p := range products {
			if p.Id == productId {
				product = &p
				break
			}
		}
		if product == nil {
			http.Error(w, "resource not found", http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(*product); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func addNewProduct(w http.ResponseWriter, r *http.Request) {
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
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler)
	r.HandleFunc("/products", getAllProducts).Methods(http.MethodGet)
	r.HandleFunc("/products", addNewProduct).Methods(http.MethodPost)
	r.HandleFunc("/products/{product_id}", getProduct).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
