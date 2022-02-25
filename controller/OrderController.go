package controller

import (
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type listOrders struct {
	Data []model.DetailedOrder `json:"data"`
}

// GetOrders by courier ID godoc
// @Summary getOrder
// @Description get orders by courier ID
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param id path int true "Courier ID"
// @Success 200 {object} model.Order
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /orders/{id} [get]
func (h *Handler) GetOrders(ctx *gin.Context) {
	var Orders []model.Order
	if err := ctx.ShouldBindJSON(&Orders); err != nil {
		logrus.Warnf("Handler getOrders (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	Orders, err = h.services.OrderApp.GetOrders(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Orders)
}

// GetOrder by order ID godoc
// @Summary getOrder
// @Description get orders by order ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} model.Order
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /order/{id} [get]
func (h *Handler) GetOrder(ctx *gin.Context) {
	var Order *model.Order
	if err := ctx.ShouldBindJSON(&Order); err != nil {
		logrus.Warnf("Handler getOrder (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	Order, err = h.services.OrderApp.GetOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Order)
}

// ChangeOrderStatus by order ID godoc
// @Summary changeOrderStatus
// @Description change order status by order ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param status query string true "status"
// @Success 200 {object} model.Order
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /order/statusChange/{id} [put]
func (h *Handler) ChangeOrderStatus(ctx *gin.Context) {
	idQuery := ctx.Param("id")
	status := ctx.Query("status")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,  model.ErrorResponse{Message: "Invalid url query"})
		return
	}
	orderId, err := h.services.OrderApp.ChangeOrderStatus(status, uint16(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Order id": orderId})
}

// GetCourierCompletedOrders by order ID godoc
// @Summary getCourierCompletedOrders
// @Description get list of completed orders by courier id
// @Tags Order
// @Accept  json
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param id query int true "idCourier"
// @Success 200 {object} listOrders
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /orders/completed [get]
func (h *Handler) GetCourierCompletedOrders(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idCourier, er := strconv.Atoi(ctx.Query("idcourier"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "expect an integer greater than 0"})
		return
	}

	DetOrders, err := h.services.OrderApp.GetCourierCompletedOrders(limit, page, idCourier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, listOrders{Data: DetOrders})

}

type listShortOrders struct {
	Data []model.Order `json:"data"`
}

// GetAllOrdersOfCourierService by order ID godoc
// @Summary getAllOrdersOfCourierService
// @Description get list of all orders by courier service id
// @Tags Order
// @Accept  json
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param id query int true "idDeliveryService"
// @Success 200 {object} listShortOrders
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /orders [get]
func (h *Handler) GetAllOrdersOfCourierService(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idService, er := strconv.Atoi(ctx.Query("iddeliveryservice"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest,  model.ErrorResponse{Message: "expect an integer greater than 0"})
		return
	}

	Orders, err := h.services.OrderApp.GetAllOrdersOfCourierService(limit, page, idService)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}

// GetCourierCompletedOrdersByMonth by order ID godoc
// @Summary getCourierCompletedOrdersByMonth
// @Description get list of completed orders by courier id sorted by month
// @Tags order
// @Accept  json
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param id query int true "idCourier"
// @Param month query int true "month"
// @Param year query int true "year"
// @Success 200 {object} listShortOrders
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /orders/byMonth [get]
func (h *Handler) GetCourierCompletedOrdersByMonth(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idCourier, er := strconv.Atoi(ctx.Query("idcourier"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest,  model.ErrorResponse{Message: "expect an integer greater than 0"})
		return
	}
	Month, er := strconv.Atoi(ctx.Query("month"))
	if er != nil || Month == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "expect an integer from 1 to 12"})
		return
	}
	Year, er := strconv.Atoi(ctx.Query("year"))
	if er != nil || Month == 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "expect an integer greater than 2021"})
		return
	}
	Orders, err := h.services.OrderApp.GetCourierCompletedOrdersByMonth(limit, page, idCourier, Month, Year)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}
