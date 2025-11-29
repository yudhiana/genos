package routes

import (
	"genos/internal/bootstrap/engine"
	"genos/modules/orders/delivery/api/rest"
)

func init() {
	application := engine.GetEngine()
	restApplication := rest.NewRestApplication()
	group := application.Group("/orders")
	group.POST("/", restApplication.CreateOrder())
}
