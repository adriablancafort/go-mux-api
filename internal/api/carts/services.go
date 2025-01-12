package carts

import (
    "github.com/adriablancafort/go-mux-api/internal/api/database"
)

func GetCartByID(id string) (*Cart, error) {
    query := `SELECT id, user_id FROM carts WHERE id = $1`
    
    var cart Cart
    if err := database.DB.QueryRow(query, id).Scan(&cart.ID, &cart.UserID); err != nil {
        return nil, err
    }

    itemsQuery := `SELECT product_id, quantity, price FROM cart_items WHERE cart_id = $1`
    rows, err := database.DB.Query(itemsQuery, id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []Item
    for rows.Next() {
        var item Item
        if err := rows.Scan(&item.ProductID, &item.Quantity, &item.Price); err != nil {
            return nil, err
        }
        items = append(items, item)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    cart.Items = items
    return &cart, nil
}

func CreateCart(cart *Cart) error {
    query := `INSERT INTO carts (id, user_id) VALUES ($1, $2)`

    _, err := database.DB.Exec(query, cart.ID, cart.UserID)
    if err != nil {
        return err
    }

    itemsQuery := `INSERT INTO cart_items (cart_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)`
    for _, item := range cart.Items {
        if _, err := database.DB.Exec(itemsQuery, cart.ID, item.ProductID, item.Quantity, item.Price); err != nil {
            return err
        }
    }

    return nil
}
