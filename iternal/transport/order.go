package transport

import (
	"L0/iternal/domain"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var order domain.Order
	if err := c.BindJSON(&order); err != nil {
		log.Println("invalid input body")
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.service.Create(context.TODO(), order.OrderUID, order); err != nil {
		log.Println("cant create order")
		newResponse(c, http.StatusInternalServerError, "cant create order")

		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"uid": order.OrderUID, 
	})
}

func (h *Handler) GetByid(c *gin.Context) {
	uid := c.Param("orderuid")
	order, err := h.service.GetById(context.TODO(), uid)
	if err != nil {
		log.Println("cant get order by id")
		newResponse(c, http.StatusBadRequest, "no order with your id")

		return
	}
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*",)
	c.JSON(http.StatusOK, order)
}

