package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /products/{id}", getProduct)

	router.Handle("/v1/", http.StripPrefix("/v1", router))

	http.ListenAndServe(":8080", router)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
		"name": "Product " + id,
	})
}