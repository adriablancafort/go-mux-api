package main

import (
    "net/http"
	"log"

    "github.com/adriablancafort/go-mux-api/internal/api/db"
    "github.com/adriablancafort/go-mux-api/internal/api/products"
)

func main() {
    db.Connect()
    defer db.Close()

    mux := http.NewServeMux()

    products.RegisterRoutes(mux)

    http.ListenAndServe(":8000", mux)
	log.Print("Server started on port 8000")
}
