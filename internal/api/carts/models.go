package carts

type Cart struct {
    ID     string  `json:"id" gorm:"primaryKey"`
    UserID string  `json:"user_id"`
    Items  []Item  `json:"items" gorm:"foreignKey:CartID"`
}

type Item struct {
    CartID    string  `json:"cart_id" gorm:"primaryKey"`
    ProductID string  `json:"product_id" gorm:"primaryKey"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}