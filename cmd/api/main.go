package main

import (
    "net/http"
	"log"

    "github.com/adriablancafort/go-mux-api/internal/api/database"
    "github.com/adriablancafort/go-mux-api/internal/api/products"
)

func main() {
    database.Connect()
    defer database.Close()

    mux := http.NewServeMux()

    products.RegisterRoutes(mux)

    http.ListenAndServe(":8000", mux)
	log.Print("Server started on port 8000")
}
