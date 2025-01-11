package products

import (
    "encoding/json"
    "net/http"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    id := r.URL.Query().Get("id")

    product := Product{
        ID:    id,
        Name:  "iPhone 12",	
        Price: 899.95,
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

    if err := json.NewEncoder(w).Encode(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}