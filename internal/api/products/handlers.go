package products

import (
	"encoding/json"
	"net/http"
    "strconv"

    "github.com/adriablancafort/go-mux-api/internal/api/authentication"
)

func RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /products", getProducts)
    mux.HandleFunc("GET /products/{id}", getProduct)
    mux.HandleFunc("POST /products", authentication.AuthMiddleware(postProduct))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    limit := 10
    if limitString := r.URL.Query().Get("limit"); limitString != "" {
        parsedLimit, err := strconv.Atoi(limitString)
        if err != nil {
            http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
            return
        }
        limit = parsedLimit
    }

    offset := 0
    if offsetString := r.URL.Query().Get("offset"); offsetString != "" {
        parsedOffset, err := strconv.Atoi(offsetString)
        if err != nil {
            http.Error(w, "Invalid offset parameter", http.StatusBadRequest)
            return
        }
        offset = parsedOffset
    }

    products, err := GetProducts(limit, offset)
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