package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetOrders(ctx *gin.Context) {
	var Orders []db.Order
	idQuery := ctx.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Orders, err = h.services.GetOrders(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such orders": err})
		return
	}
	ctx.JSON(http.StatusOK, Orders)
}

func (h *Handler) GetOrder(ctx *gin.Context) {
	var Order db.Order
	idQuery := ctx.Query("id")
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

func (h *Handler) ChangeOrderStatus(ctx *gin.Context) {
	idQuery := ctx.Query("id")
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
