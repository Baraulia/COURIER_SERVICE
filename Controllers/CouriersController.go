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

func (h *Handler) SaveCourier(c *gin.Context) {
	var Courier *db.Courier
	if err := c.ShouldBindJSON(&Courier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	Courier, err := h.services.SaveCourier(Courier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, Courier)
}
