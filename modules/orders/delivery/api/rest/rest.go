package rest

import (
	db "genos/infrastructure/sqlite"
	"genos/modules/orders/repository"
	"genos/modules/orders/usecase"

	"gorm.io/gorm"
)

type module struct {
	db *gorm.DB
	uc usecase.OrderUsecaseInterface
}

// uc := usecase.NewOrderUsecase(db.GetDatabaseConnection(), repository.NewOrderRepository(db.GetDatabaseConnection()))
func NewRestApplication() *module {
	database := db.GetDatabaseConnection()
	repo := repository.NewOrderRepository(database)
	uc := usecase.NewOrderUsecase(database, repo)
	return &module{
		db: database,
		uc: uc,
	}
}
