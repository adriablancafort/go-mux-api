package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /products/{id}", getProduct)

	http.ListenAndServe(":8080", router)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Print("Product ID: ", id, "\n")
}