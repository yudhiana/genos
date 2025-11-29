package input

// Item merepresentasikan satu produk dalam pesanan.
type Item struct {
	ProductID   uint64 `json:"product_id" binding:"required"`
	ProductName string `json:"product_name" binding:"required"`
	Qty         int    `json:"qty" binding:"required,gt=0"`
	Price       int    `json:"price" binding:"required,gt=0"`
}

// Order merepresentasikan struktur pesanan keseluruhan.
type Order struct {
	UserID uint64 `json:"user_id" binding:"required"`
	Items  []Item `json:"items" binding:"required"` // Menggunakan slice dari struct Item
}
