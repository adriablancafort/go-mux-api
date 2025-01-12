package products

import (
    "fmt"

    "github.com/adriablancafort/go-mux-api/internal/api/database"
)

func GetProducts(limit, offset int) ([]Product, error) {
    if limit < 0 || limit > 100 {
        return nil, fmt.Errorf("limit must be between 0 and 100")
    }
    if offset < 0 {
        return nil, fmt.Errorf("offset must be non-negative")
    }

    query := `SELECT id, name, price FROM products LIMIT $1 OFFSET $2`

    rows, err := database.DB.Query(query, limit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var product Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

func GetProductByID(id string) (*Product, error) {
    query := `SELECT id, name, price FROM products WHERE id = $1`
    
    var product Product
    if err := database.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price); err != nil {
        return nil, err
    }

    return &product, nil
}

func CreateProduct(product *Product) error {
    query := `INSERT INTO products (id, name, price) VALUES ($1, $2, $3)`

    _, err := database.DB.Exec(query, product.ID, product.Name, product.Price)  
    return err
}