package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetOrders(c *gin.Context) {
	var Orders []db.Order
	idQuery := c.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Orders, err = h.services.GetOrders(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, Orders)
}

func (h *Handler) GetOrder(c *gin.Context) {
	var Order []db.Order
	idQuery := c.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Order, err = h.services.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, Order)
}

func (h *Handler) ChangeOrderStatus(c *gin.Context) {
	idQuery := c.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	orderId, err := h.services.ChangeOrderStatus(uint16(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Order id": orderId})
}
