package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products/{id}", getProduct)
	mux.HandleFunc("POST /products", postProduct)

	http.ListenAndServe(":8080", mux)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
    id := r.PathValue("id")

	product := Product{
        ID:    id,
        Name:  "iPhone 12",	
        Price: 899.95,
    }

	if err := json.NewEncoder(w).Encode(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func postProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
    var product Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := json.NewEncoder(w).Encode(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}