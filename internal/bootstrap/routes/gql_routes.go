package routes

import (
	"genos/internal/bootstrap/engine"
	"genos/modules/orders/delivery/api/gql"
)

func init() {
	application := engine.GetEngine()
	group := application.Group("/gql")
	group.POST("/query", gql.GraphqlHandler())
	group.GET("/playground", gql.PlaygroundHandler())
}
