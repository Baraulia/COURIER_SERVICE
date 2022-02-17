package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Orders []db.Order

func (h *Handler) GetOrders(c *gin.Context) {

	Orders, err := h.services.GetOrders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, Orders)
}

func (h *Handler) GetOneOrder(c *gin.Context) {
	var Order []db.Order
	id := c.Query("id")
	l, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Order, err = h.services.GetOneOrder(l)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, Order)
}

func (h *Handler) ChangeOrderStatus(c *gin.Context) {
	id := c.Query("id")
	l, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	orderId, err := h.services.ChangeOrderStatus(uint16(l))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Order id": orderId})
}
