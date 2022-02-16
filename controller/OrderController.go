package controller

import (
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	_ "github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)



     var Orders []dao.Order
func (h *Handler) GetCourierCompletedOrders(c *gin.Context) {
	page,er:= strconv.Atoi(c.Query("page"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such page"})
		return
	}
	limit,er1:= strconv.Atoi(c.Query("limit"))
	if er1!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such limit"})
		return

	}
	idCourier,er:= strconv.Atoi(c.Query("idcourier"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such id"})
		return
	}
	if er==nil && page!=0 && limit!=0{
		Orders,err := h.services.OrderApp.GetCourierCompletedOrders(limit,page,idCourier)
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s",err)})
			return
		}
		c.JSON(http.StatusOK, Orders)
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s",er)})
		return
	}
}


func (h *Handler) GetAllOrdersOfCourierService(c *gin.Context) {
	page,er:= strconv.Atoi(c.Query("page"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such page"})
		return
	}
	limit,er:= strconv.Atoi(c.Query("limit"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such limit"})
		return
	}
	idService,er:= strconv.Atoi(c.Query("iddeliveryservice"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such id"})
		return
	}
	if er==nil && page!=0 && limit!=0 {
		Orders, err := h.services.OrderApp.GetAllOrdersOfCourierService( limit, page, idService)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s",err)})
			return
		}
		c.JSON(http.StatusOK, Orders)
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s",er)})
		return
	}
}

func (h *Handler) GetCourierCompletedOrdersByMonth(c *gin.Context) {
	page,er:= strconv.Atoi(c.Query("page"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such page"})
		return
	}
	limit,er:= strconv.Atoi(c.Query("limit"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such limit"})
		return
	}
	idCourier,er:= strconv.Atoi(c.Query("idcourier"))
	if er!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such id"})
		return
	}
	Month,er:= strconv.Atoi(c.Query("month"))
	if er!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no such month"})
		return
	}
	fmt.Println(limit,page,idCourier,Month)
	if er==nil && Month!=0 && page!=0 && limit!=0{
		Orders,err := h.services.OrderApp.GetCourierCompletedOrdersByMonth(	limit,page,idCourier,Month)
		if err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s",err)})
			return
		}
		c.JSON(http.StatusOK, Orders)
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s",er)})
		return
	}
}
func (h *Handler)AssigningOrderToCourier(c *gin.Context){
	/*decoder := json.NewDecoder(r.Body)
	var order dao.Order
	err := decoder.Decode(&order)
	if err != nil {
		log.Println(err)
	} */
	var order dao.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if err := h.services.AssigningOrderToCourier(order); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s",err)})
		return
	}

}