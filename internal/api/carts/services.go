package carts

import (
    "gorm.io/gorm"

    "github.com/adriablancafort/go-mux-api/internal/api/database"
)

func GetCartByID(id string) (*Cart, error) {
    var cart Cart
    result := database.DB.Preload("Items").First(&cart, "id = ?", id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &cart, nil
}

func CreateCart(cart *Cart) error {
    return database.DB.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(cart).Error; err != nil {
            return err
        }

        for i := range cart.Items {
            cart.Items[i].CartID = cart.ID
        }
        if err := tx.Create(&cart.Items).Error; err != nil {
            return err
        }

        return nil
    })
}
