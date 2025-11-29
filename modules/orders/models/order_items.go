package models

type OrderItem struct {
	ID          uint64  `gorm:"primaryKey;autoIncrement"`
	OrderID     uint64  `gorm:"not null;index"`
	ProductID   uint64  `gorm:"not null"`
	ProductName string  `gorm:"type:varchar(255);not null"`
	Qty         int     `gorm:"not null"`
	Price       float64 `gorm:"type:decimal(15,2);not null"`
}
