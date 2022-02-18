package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetOrders(c *gin.Context) {

	Orders, err := h.services.GetOrders()
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
