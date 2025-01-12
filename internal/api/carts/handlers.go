package carts

import (
    "encoding/json"
    "net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /carts/{id}", getCart)
    mux.HandleFunc("POST /carts", postCart)
}

func getCart(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    id := r.PathValue("id")

    cart, err := GetCartByID(id)
    if err != nil {
        http.Error(w, "Cart not found", http.StatusNotFound)
        return
    }

    if err := json.NewEncoder(w).Encode(cart); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func postCart(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var cart Cart
    if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := CreateCart(&cart); err != nil {
        http.Error(w, "Error creating cart", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(cart); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
