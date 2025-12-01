package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            uint64         `gorm:"primaryKey;autoIncrement"`
	UserID        uint64         `gorm:"not null;index"`
	OrderNumber   string         `gorm:"type:varchar(50);unique;not null"`
	TotalAmount   float64        `gorm:"type:decimal(15,2);not null"`
	PaymentMethod string         `gorm:"type:varchar(30);not null"`
	Status        string         `gorm:"type:varchar(20);not null;default:PENDING"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

func (o *Order) TableName() string {
	return "orders"
}
