package main

import (
	"net/http"
	"github.com/adriablancafort/go-mux-api/internal/api/products"
)

func main() {
	mux := http.NewServeMux()

	products.RegisterRoutes(mux)

	http.ListenAndServe(":8080", mux)
}
