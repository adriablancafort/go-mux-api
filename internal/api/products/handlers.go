package products

import (
    "encoding/json"
    "net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /products/", getProducts)
    mux.HandleFunc("GET /products/{id}", getProduct)
    mux.HandleFunc("POST /products", postProduct)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    products, err := GetProducts()
    if err != nil {
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(products); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func getProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    id := r.PathValue("id")

    product, err := GetProductByID(id)
    if err != nil {
        http.Error(w, "Product not found", http.StatusNotFound)
        return
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

    if err := CreateProduct(&product); err != nil {
        http.Error(w, "Error inserting product", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}