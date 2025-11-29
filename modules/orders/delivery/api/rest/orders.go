package rest

import (
	"genos/modules/orders/models/input"

	"github.com/gin-gonic/gin"
)

func (r *module) CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order input.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := r.uc.CreateOrder(c.Request.Context(), order)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message":  "User created successfully",
			"order_id": result,
		})
	}
}
