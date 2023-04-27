package transport

import (
	"L0/iternal/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var order domain.Order
	if err := c.BindJSON(&order); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.service.Create(context.TODO(), order.OrderUID, order); err != nil {
		newResponse(c, http.StatusBadRequest, "cant create order")

		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"uid": order.OrderUID, 
	})
}
