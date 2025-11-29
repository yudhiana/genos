package usecase

import (
	"context"
	commonModels "genos/internal/common/models"
	"genos/modules/orders/models"
	"genos/modules/orders/models/input"
)

type OrderUsecaseInterface interface {
	CreateOrder(ctx context.Context, input input.Order) (uint64, error)
	GetOrder(ctx context.Context, id uint64) (*models.Order, error)
	GetOrderList(ctx context.Context, option *commonModels.QueryOption) ([]models.Order, error)
}
