package controller

import (
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var Orders []dao.Order

func (h *Handler) GetCourierCompletedOrders(c *gin.Context) {
	page, er := strconv.Atoi(c.Query("page"))
	if er != nil || page == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}
	limit, er1 := strconv.Atoi(c.Query("limit"))
	if er1 != nil || limit == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}
	idCourier, er := strconv.Atoi(c.Query("idcourier"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}

	Orders, err := h.services.OrderApp.GetCourierCompletedOrders(limit, page, idCourier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, Orders)

}

func (h *Handler) GetAllOrdersOfCourierService(c *gin.Context) {
	page, er := strconv.Atoi(c.Query("page"))
	if er != nil || page == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}
	limit, er1 := strconv.Atoi(c.Query("limit"))
	if er1 != nil || limit == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}
	idService, er := strconv.Atoi(c.Query("iddeliveryservice"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}

	Orders, err := h.services.OrderApp.GetAllOrdersOfCourierService(limit, page, idService)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, Orders)

}

func (h *Handler) GetCourierCompletedOrdersByMonth(c *gin.Context) {
	page, er := strconv.Atoi(c.Query("page"))
	if er != nil || page == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}
	limit, er1 := strconv.Atoi(c.Query("limit"))
	if er1 != nil || limit == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}
	idCourier, er := strconv.Atoi(c.Query("idcourier"))
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer"})
		return
	}
	Month, er := strconv.Atoi(c.Query("month"))
	if er != nil || Month == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": " expect an integer from 1 to 12"})
		return
	}
	Orders, err := h.services.OrderApp.GetCourierCompletedOrdersByMonth(limit, page, idCourier, Month)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, Orders)

}
func (h *Handler) UpdateOrder(c *gin.Context) {
	var order dao.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if err := h.services.AssigningOrderToCourier(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}

}
