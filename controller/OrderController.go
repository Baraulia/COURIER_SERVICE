package controller

import (
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type listOrders struct {
	Data []dao.DetailedOrder `json:"data"`
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
	var Orders []dao.Order
	idQuery := ctx.Param("id")
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
	idCourier, er := strconv.Atoi(ctx.Query("idcourier"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	DetOrders, err := h.services.OrderApp.GetCourierCompletedOrders(limit, page, idCourier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listOrders{Data: DetOrders})

}

type listShortOrders struct {
	Data []dao.Order `json:"data"`
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
	idService, er := strconv.Atoi(ctx.Query("iddeliveryservice"))
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

// @Summary GetCourierCompletedOrdersByMonth
// @Description get list of completed orders by courier id sorted by month
// @Tags order
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param idcourier query int true "idcourier"
// @Param month query int true "month"
// @Param year query int true "year"
// @Success 200 {object} listShortOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders/bymonth [get]
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
	idCourier, er := strconv.Atoi(ctx.Query("idcourier"))
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

// @Summary GetDetailedOrderById
// @Description get detailed order by id
// @Tags order
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dao.DetailedOrder
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /order/detailed/{id} [get]
func (h *Handler) GetDetailedOrderById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	DetOrder, err := h.services.OrderApp.GetDetailedOrderById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, DetOrder)
}

// @Summary GetAllCompletedOrdersOfCourierService
// @Description get list of completed orders by courier service id
// @Tags order
// @Produce json
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Param iddeliveryservice query int true "iddeliveryservice"
// @Success 200 {object} listShortOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders/service/completed [get]
func (h *Handler) GetAllCompletedOrdersOfCourierService(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idService, er := strconv.Atoi(ctx.Query("iddeliveryservice"))
	if er != nil || idService <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	Orders, err := h.services.OrderApp.GetAllCompletedOrdersOfCourierService(limit, page, idService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}

// @Summary GetAllCompletedOrdersOfCourierServiceByDate
// @Description get list of completed orders by courier service id sorted bu date
// @Tags order
// @Produce json
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Param iddeliveryservice query int true "iddeliveryservice"
// @Success 200 {object} listShortOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders/service/completed/bydate [get]
func (h *Handler) GetAllCompletedOrdersOfCourierServiceByDate(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idService, er := strconv.Atoi(ctx.Query("iddeliveryservice"))
	if er != nil || idService <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	Orders, err := h.services.OrderApp.GetAllCompletedOrdersOfCourierServiceByDate(limit, page, idService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}

// @Summary GetAllCompletedOrdersOfCourierServiceByCourierId
// @Description get list of completed orders by courier service id sorted bu courier id
// @Tags order
// @Produce json
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Param iddeliveryservice query int true "iddeliveryservice"
// @Success 200 {object} listShortOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders/service/completed/bycourierid [get]
func (h *Handler) GetAllCompletedOrdersOfCourierServiceByCourierId(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idService, er := strconv.Atoi(ctx.Query("iddeliveryservice"))
	if er != nil || idService <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	Orders, err := h.services.OrderApp.GetAllCompletedOrdersOfCourierServiceByCourierId(limit, page, idService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}
