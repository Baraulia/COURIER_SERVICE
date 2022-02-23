package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"strconv"
)

// @Summary UpdateOrder
// @Tags order
// @Description assign order to courier
// @ID UpdateOrder
// @Accept  json
// @Produce json
// @Param id path int true "order_id"
// @Param input body dao.Order true "id courier"
// @Success 204
// @Failure 400 {string} string
// @Router /orders/{id} [put]
func (h *Handler) UpdateOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	var order dao.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	order.Id = id
	if err := h.services.AssigningOrderToCourier(order); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.Status(http.StatusNoContent)
}
