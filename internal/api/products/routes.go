package products

import (
    "net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /products/", getProducts)
	mux.HandleFunc("GET /products/{id}", getProduct)
	mux.HandleFunc("POST /products", postProduct)
}