package controller

import (
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type listOrders struct {
	Data []model.DetailedOrder `json:"data"`
}

// getOrders by courier ID godoc
// @Summary getOrder
// @Description get orders by courier ID
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param id path int true "Courier ID"
// @Success 200 {object} dao.Order
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /orders/{id} [get]
func (h *Handler) GetOrders(ctx *gin.Context) {
	var Orders []model.Order
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Orders, err = h.services.OrderApp.GetOrders(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such orders": err})
		return
	}
	ctx.JSON(http.StatusOK, Orders)
}

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
	var Order *model.Order
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	Order, err = h.services.OrderApp.GetOrder(id)
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
	status := ctx.Query("status")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	orderId, err := h.services.OrderApp.ChangeOrderStatus(status, uint16(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such order": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Order id": orderId})
}

// @Summary GetCourierCompletedOrders
// @Description get list of completed orders by courier id
// @Tags order
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param idcourier query int true "idcourier"
// @Success 200 {object} listOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders/completed [get]
func (h *Handler) GetCourierCompletedOrders(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	courierId, er := strconv.Atoi(ctx.Query("courierId"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	DetOrders, err := h.services.OrderApp.GetCourierCompletedOrders(limit, page, courierId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listOrders{Data: DetOrders})

}

type listShortOrders struct {
	Data []model.Order `json:"data"`
}

// @Summary GetAllOrdersOfCourierService
// @Description get list of all orders by courier service id
// @Tags order
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param iddeliveryservice query int true "iddeliveryservice"
// @Success 200 {object} listShortOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders [get]
func (h *Handler) GetAllOrdersOfCourierService(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idService, er := strconv.Atoi(ctx.Query("idDeliveryService"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	Orders, err := h.services.OrderApp.GetAllOrdersOfCourierService(limit, page, idService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}

//@Summary GetCourierCompletedOrdersByMonth
//@Description get list of completed orders by courier id sorted by month
//@Tags order
//@Produce json
//@Param page query int true "page"
//@Param limit query int true "limit"
//@Param idcourier query int true "idcourier"
//@Param month query int true "month"
//@Param year query int true "year"
//@Success 200 {object} listShortOrders
//@Failure 400 {string} string
//@Failure 500 {string} string
//@Router /orders/bymonth [get]
func (h *Handler) GetCourierCompletedOrdersByMonth(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": " limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idCourier, er := strconv.Atoi(ctx.Query("idCourier"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}
	Month, er := strconv.Atoi(ctx.Query("month"))
	if er != nil || Month == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer from 1 to 12"})
		return
	}
	Year, er := strconv.Atoi(ctx.Query("year"))
	if er != nil || Month == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 2021"})
		return
	}
	Orders, err := h.services.OrderApp.GetCourierCompletedOrdersByMonth(limit, page, idCourier, Month, Year)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}
