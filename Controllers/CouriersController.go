package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetCouriers(c *gin.Context) {
	Couriers, err := h.services.GetCouriers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, Couriers)
}

func (h *Handler) GetOneCourier(c *gin.Context) {
	var Courier db.SmallInfo
	idQuery := c.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Courier, err = h.services.GetCourier(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"No such courier": err})
		return
	}
	c.JSON(http.StatusOK, Courier)
}
