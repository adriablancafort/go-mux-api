package products

type Product struct {
    ID    string  `json:"id" gorm:"primaryKey"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}