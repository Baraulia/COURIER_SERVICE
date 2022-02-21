package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getOrders by courier ID godoc
// @Summary getOrder
// @Description get orders by courier ID
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param id path int true "Courier ID"
// @Success 200 {object} db.Order
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /orders/{id} [get]
func (h *Handler) GetOrders(c *gin.Context) {
	var Orders []db.Order
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Orders, err = h.services.GetOrders(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"No such orders": err})
		return
	}
	c.JSON(http.StatusOK, Orders)
}

// getOrder by order ID godoc
// @Summary getOrder
// @Description get orders by order ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} db.Order
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /order/{id} [get]
func (h *Handler) GetOrder(c *gin.Context) {
	var Order db.Order
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	Order, err = h.services.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"No such order": err})
		return
	}
	c.JSON(http.StatusOK, Order)
}

// putOrderStatus by order ID godoc
// @Summary putOrderStatus
// @Description put order status by order ID
// @Tags OrderStatusChange
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} db.Order
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /order/status_change/{id} [put]
func (h *Handler) ChangeOrderStatus(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	orderId, err := h.services.ChangeOrderStatus(uint16(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"No such order": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Order id": orderId})
}
