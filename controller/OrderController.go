package controller

import (
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getOrder by order ID godoc
// @Summary getOrder
// @Description get orders by order ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} dao.Order
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /order/{id} [get]
func (h *Handler) GetOrder(ctx *gin.Context) {
	var Order dao.Order
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	Order, err = h.services.GetOrder(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such order": err})
		return
	}
	ctx.JSON(http.StatusOK, Order)
}

// putOrderStatus by order ID godoc
// @Summary putOrderStatus
// @Description put order status by order ID
// @Tags OrderStatusChange
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} dao.Order
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /order/status_change/{id} [put]
func (h *Handler) ChangeOrderStatus(ctx *gin.Context) {
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	orderId, err := h.services.ChangeOrderStatus(uint16(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such order": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Order id": orderId})
}
