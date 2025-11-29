package repository

import (
	"context"

	commonModels "genos/internal/common/models"
	"genos/modules/orders/models"
	"genos/modules/orders/models/input"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (u *orderRepository) CreateOrder(ctx context.Context, input input.Order) (uint64, error) {
	db := u.db.WithContext(ctx)
	var totalAmount float64 = 0
	order := models.Order{
		UserID:      input.UserID,
		OrderNumber: "PO-" + uuid.NewString(),
		Status:      "NEW",
		OrderItems: func() (items []models.OrderItem) {
			items = make([]models.OrderItem, len(input.Items))
			for i := range input.Items {
				items[i] = models.OrderItem{
					ProductID:   input.Items[i].ProductID,
					ProductName: input.Items[i].ProductName,
					Qty:         input.Items[i].Qty,
					Price:       float64(input.Items[i].Price),
				}
				totalAmount += float64(input.Items[i].Qty) * float64(input.Items[i].Price)
			}
			return
		}(),
		TotalAmount: totalAmount,
	}
	result := db.Create(&order)
	return order.ID, result.Error
}

func (u *orderRepository) GetOrder(ctx context.Context, id uint64) (*models.Order, error) {
	db := u.db.WithContext(ctx)
	var Order models.Order
	result := db.Preload("OrderItems").First(&Order, id)
	return &Order, result.Error
}

func (u *orderRepository) GetOrderList(ctx context.Context, option *commonModels.QueryOption) ([]models.Order, error) {
	db := u.db.WithContext(ctx)
	var Orders []models.Order
	var limit int = 10
	var offset int = 0
	var order string = "id DESC"

	if option != nil {
		// Apply pagination if provided
		limit = option.Limit
		if option.Order != "" {
			order = option.Order
		}
		offset = (option.Page - 1) * limit
	}

	query := db.Model(&models.Order{}).Order(order)
	query = query.Limit(limit).Offset(offset)
	result := query.Find(&Orders)
	return Orders, result.Error
}

func (u *orderRepository) GetOrderByUserID(ctx context.Context, userID uint64, option *commonModels.QueryOption) ([]models.Order, error) {
	db := u.db.WithContext(ctx)
	var Orders []models.Order
	var limit int = 10
	var offset int = 0
	var order string = "id DESC"

	if option != nil {
		// Apply pagination if provided
		limit = option.Limit
		if option.Order != "" {
			order = option.Order
		}
		offset = (option.Page - 1) * limit
	}

	query := db.Model(&models.Order{}).Where("user_id = ?", userID).Order(order)
	query = query.Limit(limit).Offset(offset)
	result := query.Preload("OrderItems").Find(&Orders)
	return Orders, result.Error
}
