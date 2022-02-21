package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"strconv"
)


// @Summary UpdateOrder
// @Tags order
// @Description assign order to courier
// @ID UpdateOrder
// @Accept  json
// @Produce json
// @Param input body dao.Order true "id courier/id order"
// @Success 204
// @Failure 400,404 {string} string
// @Failure 500 {string} string
// @Router /orders [put]
func (h *Handler) UpdateOrder(ctx *gin.Context) {
	var order dao.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if err := h.services.AssigningOrderToCourier(order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.Status(http.StatusNoContent)
}
type  listOrders struct {
	Data []dao.DetailedOrder `json:"data"`
}
// @Summary GetAllCompletedOrdersByService
// @Description get list of orderss
// @Tags order
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param idservice query int true "idservice"
// @Success 200 {object} listOrders
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders/completed [get]
func (h *Handler) GetAllCompletedOrdersByService(ctx *gin.Context){
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
	idService, er := strconv.Atoi(ctx.Query("idservice"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	DetOrders, err := h.services.OrderApp.GetAllServiceCompletedOrders(limit, page, idService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listOrders{Data: DetOrders})
}

// @Summary GetDetailedOrdersById
// @Description get detailed order by id
// @Tags order
// @Produce json
// @Param id query int true "id"
// @Success 200 {object} dao.DetailedOrder
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /orders [get]
func(h *Handler)  GetDetailedOrdersById(ctx *gin.Context){
	Id, er := strconv.Atoi(ctx.Query("id"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	DetOrder, err := h.services.OrderApp.GetDetailedOrdersById(Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, DetOrder)
}
