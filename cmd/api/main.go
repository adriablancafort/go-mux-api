package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products/{id}", getProduct)

	http.ListenAndServe(":8080", mux)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    w.Header().Set("Content-Type", "application/json")

    product := map[string]interface{}{
        "id":   id,
        "name": "Product " + id,
    }

	err := json.NewEncoder(w).Encode(product); 
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}