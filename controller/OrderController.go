package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"strconv"
)

type listOrders struct {
	Data []dao.DetailedOrder `json:"data"`
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
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
// @Success 200 {object} listShortOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders/bymonth [get]
func (h *Handler) GetCourierCompletedOrdersByMonth(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer greater than 0"})
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
	Orders, err := h.services.OrderApp.GetCourierCompletedOrdersByMonth(limit, page, idCourier, Month)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listShortOrders{Data: Orders})

}
