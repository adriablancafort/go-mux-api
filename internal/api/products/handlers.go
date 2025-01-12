package products

import (
    "encoding/json"
    "net/http"
    "log"

    "github.com/adriablancafort/go-mux-api/internal/api/database"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var products []Product
    query := `SELECT id, name, price FROM products`
    rows, err := database.DB.Query(query)
    if err != nil {
        log.Println("Error fetching products:", err)
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var product Product
        err := rows.Scan(&product.ID, &product.Name, &product.Price)
        if err != nil {
            log.Println("Error scanning product:", err)
            http.Error(w, "Error scanning product", http.StatusInternalServerError)
            return
        }
        products = append(products, product)
    }

    if err := rows.Err(); err != nil {
        log.Println("Error iterating over rows:", err)
        http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
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
    if id == "" {
        http.Error(w, "Missing product ID", http.StatusBadRequest)
        return
    }

    var err error
    var product Product
    query := `SELECT id, name, price FROM products WHERE id = $1`
    row := database.DB.QueryRow(query, id)
    err = row.Scan(&product.ID, &product.Name, &product.Price)
    if err != nil {
        log.Println("Error fetching product:", err)
        http.Error(w, "Product not found", http.StatusNotFound)
        return
    }

    err = json.NewEncoder(w).Encode(product)
    if err != nil {
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

    query := `INSERT INTO products (id, name, price) VALUES ($1, $2, $3)`
    _, err := database.DB.Exec(query, product.ID, product.Name, product.Price)
    if err != nil {
        log.Println("Error inserting product:", err)
        http.Error(w, "Error inserting product", http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}