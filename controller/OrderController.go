package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
)

func (h *Handler) UpdateOrder(c *gin.Context) {
	var order dao.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if err := h.services.AssigningOrderToCourier(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
}
