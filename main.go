package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /products/{id}", getProduct)

	http.ListenAndServe(":8080", router)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
		"title": "Product " + id,
	})
}