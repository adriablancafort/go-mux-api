package carts

type Cart struct {
    ID     string  `json:"id"`
    UserID string  `json:"user_id"`
    Items  []Item  `json:"items"`
}

type Item struct {
    ProductID string  `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}