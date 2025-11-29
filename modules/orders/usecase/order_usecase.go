package usecase

import (
	"context"

	commonModels "genos/internal/common/models"
	"genos/modules/orders/models"
	"genos/modules/orders/models/input"
	"genos/modules/orders/repository"

	"gorm.io/gorm"
)

type orderUsecase struct {
	db   *gorm.DB
	repo repository.OrderRepository
}

func NewOrderUsecase(db *gorm.DB, repo repository.OrderRepository) OrderUsecaseInterface {
	return &orderUsecase{
		db:   db,
		repo: repo,
	}
}

func (uc *orderUsecase) CreateOrder(ctx context.Context, input input.Order) (id uint64, err error) {
	id, err = uc.repo.CreateOrder(ctx, input)
	if err != nil {
		return
	}
	return id, nil
}

func (uc *orderUsecase) GetOrder(ctx context.Context, id uint64) (*models.Order, error) {
	orderData, err := uc.repo.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	return orderData, nil
}

func (uc *orderUsecase) GetOrderList(ctx context.Context, option *commonModels.QueryOption) ([]models.Order, error) {
	orderList, err := uc.repo.GetOrderList(ctx, option)
	if err != nil {
		return nil, err
	}

	return orderList, nil
}

func (uc *orderUsecase) GetOrderByUserID(ctx context.Context, userID uint64, option *commonModels.QueryOption) ([]models.Order, error) {
	orderList, err := uc.repo.GetOrderByUserID(ctx, userID, option)
	if err != nil {
		return nil, err
	}

	return orderList, nil
}
