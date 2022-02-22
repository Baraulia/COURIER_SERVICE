package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
)


// @Summary UpdateOrder
// @Tags order
// @Description assign order to courier
// @ID UpdateOrder
// @Accept  json
// @Produce json
// @Param input body dao.Order true "id courier/id order"
// @Success 204
// @Failure 400,404 {string} string
// @Failure 500 {string} string
// @Router /orders [put]
func (h *Handler) UpdateOrder(ctx *gin.Context) {
	var order dao.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if err := h.services.AssigningOrderToCourier(order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.Status(http.StatusNoContent)
}
