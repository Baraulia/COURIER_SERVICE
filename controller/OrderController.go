package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"strconv"
)

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
}

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
	ctx.JSON(http.StatusOK, DetOrders)
}

func(h *Handler)  GetDetailedOrdersById(ctx *gin.Context){
	idOrder, er := strconv.Atoi(ctx.Query("idorder"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	DetOrder, err := h.services.OrderApp.GetDetailedOrdersById(idOrder)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, DetOrder)
}
