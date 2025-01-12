package main

import (
    "net/http"
	"log"

	"github.com/joho/godotenv"

    "github.com/adriablancafort/go-mux-api/internal/api/database"
    "github.com/adriablancafort/go-mux-api/internal/api/products"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println(".env file not found")
    }

    database.Connect()
    defer database.Close()

    mux := http.NewServeMux()

	mux.Handle("/v1/", http.StripPrefix("/v1", mux))
	
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    })
	
	products.RegisterRoutes(mux)

	log.Print("Server started on port 8000")
    http.ListenAndServe(":8000", mux)
}
