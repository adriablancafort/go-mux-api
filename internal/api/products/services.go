package products

import (
    "github.com/adriablancafort/go-mux-api/internal/api/database"
)

func GetProducts(limit, offset int) ([]Product, error) {
    var products []Product
    result := database.DB.Limit(limit).Offset(offset).Find(&products)
    return products, result.Error
}

func GetProductByID(id string) (*Product, error) {
    var product Product
    result := database.DB.First(&product, "id = ?", id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &product, nil
}

func CreateProduct(product *Product) error {
    result := database.DB.Create(product)
    return result.Error
}