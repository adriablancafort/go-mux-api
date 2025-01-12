package products

import (
    "github.com/adriablancafort/go-mux-api/internal/api/database"
)

func GetProducts() ([]Product, error) {
    query := `SELECT id, name, price FROM products`
    rows, err := database.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var product Product
        err := rows.Scan(&product.ID, &product.Name, &product.Price)
        if err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    err = rows.Err(); 
    if err != nil {
        return nil, err
    }

    return products, nil
}

func GetProductByID(id string) (*Product, error) {
    query := `SELECT id, name, price FROM products WHERE id = $1`
    
    var product Product
    err := database.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price); 
    if err != nil {
        return nil, err
    }

    return &product, nil
}

func CreateProduct(product *Product) error {
    query := `INSERT INTO products (id, name, price) VALUES ($1, $2, $3)`

    _, err := database.DB.Exec(query, product.ID, product.Name, product.Price)  
    return err
}